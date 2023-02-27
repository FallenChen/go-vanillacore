package io

type VirtualChannel struct {
	FileChannel Channel
}

func NewVirtualChannel(fileChannel *Channel) VirtualChannel {

	return VirtualChannel{}
}

func (vir VirtualChannel) Read(buf Buffer, position int64) error {

	return nil
}

func (vir VirtualChannel) Write(buf Buffer, position int64) error {

	return nil
}

func (vir VirtualChannel) Append(buf Buffer) (int64, error) {

	return 0, nil
}

func (vir VirtualChannel) Size() (int64, error) {
	return 0, nil
}

func (vir VirtualChannel) Close() error {

	return nil
}
