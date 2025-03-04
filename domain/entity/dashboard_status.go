package entity

type DashboardStatus int

const (
	DashboardStatusStarter            DashboardStatus = -1
	DashboardStatusAuthPending        DashboardStatus = 0
	DashboardStatusAuthErrorInstagram DashboardStatus = 1
	DashboardStatusAuthErrorWordpress DashboardStatus = 2
	DashboardStatusAuthSuccess        DashboardStatus = 3
	DashboardStatusTokenExpired       DashboardStatus = 4
	DashboardStatusExecuteSuccess     DashboardStatus = 5
	DashboardStatusModStartDate       DashboardStatus = 6
	DashboardStatusUpdateInformation  DashboardStatus = 7
)

func (d DashboardStatus) ToPrompt() string {
	switch d {
	case DashboardStatusStarter:
		return starter()
	case DashboardStatusAuthPending:
		return authPending()
	case DashboardStatusAuthErrorInstagram:
		return authErrorInstagram()
	case DashboardStatusAuthErrorWordpress:
		return authErrorWordpress()
	case DashboardStatusAuthSuccess:
		return authSuccess()
	case DashboardStatusTokenExpired:
		return tokenExpired()
	case DashboardStatusExecuteSuccess:
		return executeSuccess()
	case DashboardStatusModStartDate:
		return modStartDate()
	case DashboardStatusUpdateInformation:
		return ""
	default:
		return ""
	}
}

func starter() string {
	return "ユーザーに初めましての挨拶をお願いします。そのあと、Instagramとの連携をしてもらってください。手順は以下の通りです\n" + getFacebookAuthentication()
}

func authPending() string {
	return "ユーザーに挨拶をし、まだユーザーにInstagramとの連携できていないことを伝えてください。手順は以下の通りです\n" + getFacebookAuthentication()
}

func tokenExpired() string {
	return "ユーザーに挨拶をしてください。\n" +
		"そして、トークンの有効期限が切れてしまっているのでInstagramとの連携を再度してもらってください。\n" +
		"手順は、まず「Instagramとつなぐ」ボタンをクリックします。そして「再リンク」をクリックします。これで終わりです。"
}

func authErrorInstagram() string {
	return `ユーザーがインスタグラムとの連携に失敗してしまいました。
「ビジネス」->「ページ」->「インスタグラム」の選択が間違っていると考えられます。
以下の内容をもう一度確認してもらってください。
・「sd-a-rootがアクセスするビジネスを選択」にて、「現在および今後のビジネスすべてにオプトイン」を選択して「続行」ボタンを押します。
・「sd-a-rootがアクセスするページを選択」にて、「現在および今後のページすべてにオプトイン」を選択して「続行」ボタンを押します。
・「sd-a-rootがアクセスするInstagramアカウントを選択」にて、「現在のInstagramアカウントのみにオプトイン」を選択し、今回連携させたいInstagramアカウントを一つだけ選択して、「続行」ボタンを押します。
「sd-a-rootがアクセスするInstagramアカウントを選択」の時に連携したいInstagramアカウントが出てこない場合、ログインしているFacebookアカウントが間違っている可能性があります。`
}

func authErrorWordpress() string {
	return `ユーザーがWordPressとの連携に失敗してしまいました。
管理者側の対応が漏れているので、ユーザーにはお待ちいただいてください。
エラーが出ていることは開発者側には通知が行っているので別途連絡は大丈夫ですが、お急ぎの場合は、直接ご連絡くださいとご案内お願いします。`
}

func authSuccess() string {
	return "連携がうまくいったとお伝えください。一緒に喜んでください\n" +
		"念のため、連携したインスタグラムのアカウントが本当に意図したアカウントのものか、もう一度確認してもらってください。"
}

func getFacebookAuthentication() string {
	return `・「Instagramとつなぐ」ボタンをクリックします。すると、ダイアログが表示されます。
・まず、Facebookアカウントでログイン処理をします。
・「sd-a-rootがアクセスするビジネスを選択」にて、「現在および今後のビジネスすべてにオプトイン」を選択して「続行」ボタンを押します。
・「sd-a-rootがアクセスするページを選択」にて、「現在および今後のページすべてにオプトイン」を選択して「続行」ボタンを押します。
・「sd-a-rootがアクセスするInstagramアカウントを選択」にて、「現在のInstagramアカウントのみにオプトイン」を選択し、今回連携させたいInstagramアカウントを一つだけ選択して、「続行」ボタンを押します。
・「保存」をクリックします。
・「●●はsd-a-rootにリンクされています」と表示されたら、「OK」ボタンをクリックします。
・連携が完了します。`
}

func modStartDate() string {
	return "連携日時を変更がうまくいったことを伝えてください。一度、連携したものは対象外になるので日時が過去の連携日と重なっても大丈夫だと伝えてください。"
}

func executeSuccess() string {
	return "連携がうまくいったことを伝えて、一緒に喜んでください"
}
