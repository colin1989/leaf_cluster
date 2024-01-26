package gate

func encrypt(data []byte) {

	size := len(data)

	for i := 0; i < size; i++ {

		data[i] = byte(int(data[i]) ^ (size - i))

	}

}
