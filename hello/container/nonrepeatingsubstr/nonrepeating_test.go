package main

import "testing"

//可在命令行执行 go test .		查看测试结果
//可在命令行执行 go test -coverprofile=c.out 		查看测试的代码覆盖率
//c.out 是一个文件 可以用 less c.out 查看
//也可执行命令行 go tool cover -html=c.out 在html页面中查看
func Test_lenthOfNonRepeatingSubStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name          string
		args          args
		wantMaxLength int
	}{
		// TODO: Add test cases.
		{"normal cases", args{"abcabcbb"}, 3},
		{"edge cases", args{""}, 0},
		{"chinese support cases", args{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxLength := lenthOfNonRepeatingSubStr(tt.args.s); gotMaxLength != tt.wantMaxLength {
				t.Errorf("lenthOfNonRepeatingSubStr() = %v, want %v", gotMaxLength, tt.wantMaxLength)
			}
		})
	}
}

//性能测试
//也可以在命令行 输入 go test -bench . 进行测试
//命令行输入 go test -bench . -cpuprofile cpu.out   用go提供的工具来查看代码那里用时间最多
//go tool pprof cpu.out		//查看 cpu.out这个二进制文件
//输入 go tool pprof cpu.out 后可以终端进行交互，
//交互输入 help 可以看我们能做什么事
//输入	web  可以打开一个网页  这个网页需要安装 gvedit 不然会报错 failed to execute dot. Is Graphviz installed? Error: exec: "dot": executable file not found in $PATH
//通过 homebrew 安装 brew install graphviz  安装后需重新启动终端
//输入 exit or quit 可以退出交互
func Benchmark_lenthOfNonRepeatingSubStr(b *testing.B) {
	//性能测试一般选最难的数据来做，做多次
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"

	//增加字符串的长度，看看时间占比怎么样
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	wantMaxLength := 8
	//b.N 性能测试要测试的次数，不用人为去写，go有算法
	//上面是准备数据的时间，不需要计算在内
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if gotMaxLength := lenthOfNonRepeatingSubStr(s); gotMaxLength != wantMaxLength {
			b.Errorf("lenthOfNonRepeatingSubStr() = %v, want %v", gotMaxLength, wantMaxLength)
		}
	}

}
