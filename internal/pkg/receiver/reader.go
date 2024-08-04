package receiver

type TReader struct {
	ID         uint
	Serial     string
	Worker     string
	WorkerCard string
}

func (r *TReader) SetDeviceId(serial string, new_id int) {

}

func (r *TReader) DeleteData(data []byte) {
	return
}

func (r *TReader) GetSerial() string {
	return ""
}

func (r *TReader) Beep(time uint, skip uint) {
	return
}

func (r *TReader) TextOut(col int, row int, text string) {
	return
}

func (r *TReader) LogOn(worker string) {
	return
}

func (r *TReader) LogOff() {
	return
}
