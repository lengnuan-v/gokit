// +----------------------------------------------------------------------
// | tts 语音合成
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月05日
// +----------------------------------------------------------------------

package gokit

import (
	"bytes"
	"encoding/json"
	gokit "github.com/lengnuan-v/gokit/utils"
	"io/ioutil"
	"net/http"
)

const ttsApiUrl = "https://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/tts"

type TTS struct {
	Appkey     string // Appkey
	Text       string // 待合成的文本
	Token      string // 服务鉴权Token
	Voice      string // 发音人，默认是xiaoyun
	Format     string // 音频编码格式，支持的格式：pcm、wav、mp3，默认是pcm
	SampleRate string // 音频采样率，支持16000Hz、8000Hz，默认是16000Hz
	Volume     int    // volume 音量，范围是0~100，可选，默认50
	SpeechRate int    // SpeechRate 语速，范围是-500~500，可选，默认是0
	PitchRate  int    // PitchRate 语调，范围是-500~500，可选，默认是0
}

func (t *TTS) GetAliyunTts() ([]byte, error) {
	var err error
	bodyContent := make(map[string]interface{})
	bodyContent["appkey"] = t.Appkey
	bodyContent["text"] = t.Text
	bodyContent["token"] = t.Token
	bodyContent["format"] = t.Format
	bodyContent["sample_rate"] = t.SampleRate
	bodyContent["speech_rate"] = t.SpeechRate
	bodyContent["pitch_rate"] = t.PitchRate
	// 音频编码格式如果不填写就使用默认
	if gokit.IsEmpty(bodyContent["format"]) == true {
		bodyContent["format"] = "mp3"
	}
	bodyContent["voice"] = t.Voice
	// 发音人如果不填写就使用默认
	if gokit.IsEmpty(bodyContent["voice"]) == true {
		bodyContent["voice"] = "siyue"
	}
	bodyContent["volume"] = t.Volume
	// 音量如果不填写就使用默认
	if gokit.IsEmpty(bodyContent["volume"]) == true {
		bodyContent["volume"] = 100
	}
	var bodyJson []byte
	if bodyJson, err = json.Marshal(bodyContent); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = http.Post(ttsApiUrl, "application/json;charset=utf-8", bytes.NewBuffer([]byte(bodyJson))); err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if response.Header.Get("Content-Type") == "audio/mpeg" && err == nil {
		return []byte(body), nil
	} else {
		return nil, err
	}
}