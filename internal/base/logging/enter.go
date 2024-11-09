package logging

//
//import (
//	"gopkg.in/natefinch/lumberjack.v2"
//	"io"
//	"log/slog"
//	"os"
//)
//
//func init() {
//	once.Do(func() {
//		var logOut io.Writer
//		logOut = os.Stdout
//		switch logType {
//		default:
//			slog.Info("Unknown Log Output Type")
//		case LogTypeStdout:
//		case LogTypeFile:
//			if rolling {
//				logOut = getAsyncFileIoRolling()
//			} else {
//				logOut = getFileIo()
//			}
//		}
//		logLevel := slog.LevelInfo
//		if debug {
//			logLevel = slog.LevelDebug
//		}
//		slog.SetDefault(slog.New(slog.NewJSONHandler(logOut, &slog.HandlerOptions{
//			AddSource:   true,
//			ReplaceAttr: replace,
//			Level:       logLevel,
//		})))
//	})
//}
//
//func getAsyncFileIoRolling() *AsyncW {
//	aw = AsyncLumberjack(&lumberjack.Logger{
//		Filename:   logPath,
//		MaxSize:    maxSize,    // megabytes
//		MaxBackups: maxBackUps, // maxBackUps
//		MaxAge:     maxAge,     //days
//		LocalTime:  true,
//		Compress:   false, // disabled by default
//	})
//	return aw
//}
//
//func Shutdown() {
//	aw.Stop()
//	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
//		AddSource:   true,
//		ReplaceAttr: replace,
//	})))
//	slog.Info("logging 👋")
//}
