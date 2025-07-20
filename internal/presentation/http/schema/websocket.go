package schema

type EventFromUserSchema struct {
	EventType string  `json:"event_type"`          //イベントの種類
	CardID    string  `json:"card_id,omitempty"`   //ユーザーが持っている電話カードのID
	StringID  string  `json:"string_id,omitempty"` //ユーザーが持っている糸のID
	Distance  float64 `json:"distance,omitempty"`  //ユーザー間の距離
}

type EventTypeFromUser int

const (
	EventTypeUnknown             EventTypeFromUser = iota
	EventTypeCallRequest                           //電話をかけるイベント(糸、テレフォンカード、電話番号の指定)
	EventTypeCallAccept                            //電話を受けるイベント
	EventTypeConnected                             //通話が開始されたイベント
	EventTypeChangeTerephoneCard                   //使用中のテレフォンカードの変更イベント
	EventTypeChangeString                          //使用中の糸の変更イベント
	EventTypeErrorFromUser                         //エラーイベント
)

type EventTypeFromServer int

const (
	EventTypeUnknownFromServer EventTypeFromServer = iota
	EventTypeNotifycation                          //着信イベント
	EventTypeCallReject                            //通話開始を拒否するイベント
	EventTypeCallEnd                               //通話を終了するイベント
	EventTypeStartNearBy                           //近くのユーザーに電話をかけるイベント
	EventTypeBalanceRunOut                         //通話中のテレフォンカードの残高がなくなったイベント(10sの猶予のあとEventTypeCallEndが送られる)
	EventTypeCallingResult                         //通話中のイベント(糸の耐久値, テレフォンカードの残高)
	EventTypeErrorFromServer                       //エラーイベント
)

type EventFromServerSchema struct {
	EventType  EventTypeFromServer `json:"event_type"`           //イベントの種類
	Durability int                 `json:"durability,omitempty"` //使用中の糸の耐久値
	Balance    int                 `json:"balance,omitempty"`    //使用中のテレフォンカードの残高
}
