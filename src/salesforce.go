package main

import (
	"github.com/sclevine/agouti"

	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	Sleep_TIME = 5 * time.Second

	SalesforceLoginUrl = `https://login.salesforce.com/`

	// Html Attribute Name In Login Menu
	UserNameID = "username"
	PasswordID = "password"
	LoginID    = "Login"

	// Html Attribute Name In Main Menu
	WorkTabID       = "01r2800000085Bg_Tab"
	MonthListID     = "yearMonthList"
	WorkTdID        = "ttvTimeSt"
	ProjectButtonID = "dailyWorkCell"

	// Html Attribute Name In Dialog Input Time
	StartTimeID          = "startTime"
	EndTimeID            = "endTime"
	RegisterWorkButtonID = "dlgInpTimeOk"

	// Html Attribute Name In Dialog Input Projects
	WorkTableID             = "empWorkTableBody"
	WorkTRCssClass1         = "tr.odd"
	WorkTRCssClass2         = "tr.even"
	ProjectClass            = "td > div.name"
	ProjectInputClass       = "inputime"
	RegisterProjectButtonID = "empWorkOk"

	MonthListTextFormat = "2006年01月"
	WorkTdDateFormat    = "2006-01-02"
	InputTimeFormat     = "15:04"

	BackSpaceKey = "\ue003"
)

func (sf *salesforce) inputWorkDuration(start, end string) error {

	// 稼働時間を入力
	err := sf.Page.FindByID(StartTimeID).Fill(start)
	if err != nil {
		return err
	}

	err = sf.Page.FindByID(EndTimeID).Fill(end)
	if err != nil {
		return err
	}

	// 登録
	err = sf.Page.FindByID(RegisterWorkButtonID).Click()
	if err != nil {
		return err
	}

	return nil
}

type projectRepo struct {
	Name     string
	CssClass string
	Number   int
}

func (sf *salesforce) newProjectRepoByCssClass(count int, cssClass string) ([]projectRepo, error) {
	var repo []projectRepo
	for i := 0; i < count; i++ {
		text, err := sf.Page.FindByID(WorkTableID).All(cssClass + " > " + ProjectClass).At(i).Text()
		if err != nil {
			return repo, err
		}
		repo = append(repo, projectRepo{
			Name:     text,
			CssClass: cssClass,
			Number:   i,
		})
	}
	return repo, nil
}

func (sf *salesforce) newProjectsRepo() ([]projectRepo, error) {
	var repo []projectRepo

	//cssクラス毎にセレクターを取得し、登録されているプロジェクトのリストを作成する
	for _, class := range []string{WorkTRCssClass1, WorkTRCssClass2} {
		class_count, err := sf.Page.FindByID(WorkTableID).All(class).Count()
		if err != err {
			return repo, err
		}
		classRepo, err := sf.newProjectRepoByCssClass(class_count, class)
		if err != nil {
			return repo, err
		}
		repo = append(repo, classRepo...)
	}

	return repo, nil
}

func findProject(project Project, repo []projectRepo) (bool, projectRepo) {
	re := regexp.MustCompile(project.Name)
	for _, v := range repo {
		if re.MatchString(v.Name) {
			return true, v
		}
	}
	return false, projectRepo{}
}

func (sf *salesforce) inputProjectsDuration(projects []Project) error {

	// プロジェクトリストを取得
	repo, err := sf.newProjectsRepo()
	if err != nil {
		return err
	}
	for _, inputProject := range projects {
		exists, project := findProject(inputProject, repo)
		if exists == false {
			return errors.New(fmt.Sprintf("%s is not exists.", inputProject.Name))
		}

		// 工数入力
		err = sf.Page.FindByID(WorkTableID).All(project.CssClass).At(project.Number).FindByClass(
			ProjectInputClass).Fill(
			strings.Repeat(BackSpaceKey, 5) +
				inputProject.Duration)
		if err != nil {
			return err
		}
	}

	// 登録
	if err = sf.Page.FindByID(RegisterProjectButtonID).Click(); err != nil {
		return err
	}
	return nil
}

type account struct {
	UserName string
	Password string
}

type salesforce struct {
	Account account
	Page    *agouti.Page
}

func (d *Driver) NewSalesForce(username, password string) (*salesforce, error) {
	page, err := d.NewPage()
	if err != nil {
		return nil, err
	}
	if err := page.Navigate(SalesforceLoginUrl); err != nil {
		return nil, err
	}
	return &salesforce{
		Account: account{
			UserName: username,
			Password: password,
		},
		Page: page,
	}, nil
}

func (sf *salesforce) Login() error {
	// ID, Passの要素を取得し、値を設定
	sf.Page.FindByID(UserNameID).Fill(sf.Account.UserName)
	sf.Page.FindByID(PasswordID).Fill(sf.Account.Password)

	// formをサブミット
	if err := sf.Page.FindByID(LoginID).Submit(); err != nil {
		return err
	}

	time.Sleep(Sleep_TIME)
	return nil

}

func (sf *salesforce) RegisterWork(dailywork DailyWork) error {

	// 勤務表タブをクリック
	if err := sf.Page.FindByID(WorkTabID).Click(); err != nil {
		return err
	}

	// ちょっと待つ
	time.Sleep(Sleep_TIME)

	// 月を選択
	err := sf.Page.FindByID(MonthListID).Select(
		dailywork.TypeChange.Date.Format(MonthListTextFormat))
	if err != nil {
		return err
	}

	// ちょっと待つ
	time.Sleep(Sleep_TIME)

	// 「出社」項目をクリック
	err = sf.Page.FindByID(
		WorkTdID +
			dailywork.TypeChange.Date.Format(WorkTdDateFormat)).Click()
	if err != nil {
		return err
	}

	// ちょっと待つ
	time.Sleep(Sleep_TIME)

	err = sf.inputWorkDuration(dailywork.Start, dailywork.End)
	if err != nil {
		return err
	}

	// ちょっと待つ
	time.Sleep(Sleep_TIME)

	// 「工数」項目をクリック
	err = sf.Page.FindByID(
		ProjectButtonID +
			dailywork.TypeChange.Date.Format(WorkTdDateFormat)).Click()
	if err != nil {
		return err
	}

	// ちょっと待つ
	time.Sleep(Sleep_TIME)

	err = sf.inputProjectsDuration(dailywork.Projects)
	if err != nil {
		return err
	}

	return nil
}
