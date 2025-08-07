package helper

// 六个chen通过函数生成一个chen的函数
// 例如：
//	add := helper.Operate3(ac, bc, cc, func(a, b, c int) int {
//	  return a + b + c
//	})
func Operate7[A any, B any, C any, D any, E any, F any, G any, R any](ac <-chan A, bc <-chan B, cc <-chan C, dc <-chan D, ec <-chan E, fc <-chan F, gc <-chan G, o func(A, B, C, D, E, F, G) R) <-chan R {
	// 创建一个输出通道
	rc := make(chan R)

	go func() {
		// 自动关闭输出通道
		defer close(rc)

		for {
			an, ok := <-ac
			if !ok {
				break
			}

			bn, ok := <-bc
			if !ok {
				break
			}

			cn, ok := <-cc
			if !ok {
				break
			}

			dn, ok := <-dc
			if !ok {
				break
			}

			en, ok := <-ec
			if !ok {
				break
			}

			fn, ok := <-fc
			if !ok {
				break
			}
			gn, ok := <-gc
			if !ok {
				break
			}
			// 任意一个通道没有数据，就退出循环

			rc <- o(an, bn, cn, dn, en, fn, gn)
		}

		// 排空所有通道
		Drain(ac)
		Drain(bc)
		Drain(cc)
		Drain(dc)
		Drain(ec)
		Drain(fc)
		Drain(gc)
	}()

	return rc
}
