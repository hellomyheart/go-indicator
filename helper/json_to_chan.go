package helper

import (
	"encoding/json"
	"io"
	"log/slog"
)

// JSONToChan 以JSON格式从指定读取器读取值到值通道。
func JSONToChan[T any](r io.Reader) <-chan T {
	return JSONToChanWithLogger[T](r, slog.Default())
}

// JSONToChanWithLogger 以JSON格式从指定读取器读取值到值通道。增加日志输出配置
func JSONToChanWithLogger[T any](r io.Reader, logger *slog.Logger) <-chan T {
	// 创建一个chan
	c := make(chan T)

	go func() {
		defer close(c)

		// 解析json
		decoder := json.NewDecoder(r)

		// 获取Token
		token, err := decoder.Token()
		if err != nil {
			logger.Error("Unable to read token.", "error", err)
			return
		}

		// 如果不是数组
		if token != json.Delim('[') {
			logger.Error("Expecting start of array.", "token", token)
			return
		}

		// 遍历json数组
		for decoder.More() {
			var value T

			// 解析json
			err = decoder.Decode(&value)
			if err != nil {
				logger.Error("Unable to decode value.", "error", err)
				return
			}

			// 发送json
			c <- value
		}

		//数组结束判断
		token, err = decoder.Token()
		if err != nil {
			logger.Error("Unable to read token.", "error", err)
			return
		}

		if token != json.Delim(']') {
			logger.Error("Expecting end of array.", "token", token)
			return
		}
	}()

	return c
}
