package errs

// 内部エラー型(データベース接続エラーなど)
type CRUDError struct {
	message string // エラーメッセージ
}

// エラーメッセージを返すメソッド
func (e *CRUDError) Error() string {
	return e.message
}

// コンストラクタ
func NewCRUDError(message string) *CRUDError {
	return &CRUDError{message: message}
}
