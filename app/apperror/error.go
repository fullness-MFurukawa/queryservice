package apperror

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// *****
// アプリケーションエラー型
// *****
type AppError struct {
	status  string // ステータスHTTPステータスを利用する
	message string // エラーメッセージ
}

/*
エラーメッセージを返す
*/
func (e *AppError) Error() string {
	// ステータスとメッセージを組み合わせた文字列を返す
	return fmt.Sprintf("status:%s,message:%s", e.status, e.message)
}

/*
コンストラクタ
*/
func NewAppError(status string, message string) error {
	return &AppError{status: status, message: message}
}

/*
コンストラクタ(内部エラー)
*/
func NewInternalError(err error) error {
	log.Error().Int("Status", 500).Msg(err.Error())
	return NewAppError("500", "現在サービスの提供ができません。")
}
