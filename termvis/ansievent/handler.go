package ansievent

type Handler struct {
	Events *[]interface{}
}

func (s *Handler) Print(b byte) error {
	*s.Events = append(*s.Events, Print{b})
	return nil
}

func (s *Handler) Execute(b byte) error {
	*s.Events = append(*s.Events, Execute{b})
	return nil
}

func (s *Handler) CUU(param int) error {
	*s.Events = append(*s.Events, CUU{param})
	return nil
}

func (s *Handler) CUD(param int) error {
	*s.Events = append(*s.Events, CUD{param})
	return nil
}

func (s *Handler) CUF(param int) error {
	*s.Events = append(*s.Events, CUF{param})
	return nil
}

func (s *Handler) CUB(param int) error {
	*s.Events = append(*s.Events, CUB{param})
	return nil
}

func (s *Handler) CNL(param int) error {
	*s.Events = append(*s.Events, CNL{param})
	return nil
}

func (s *Handler) CPL(param int) error {
	*s.Events = append(*s.Events, CPL{param})
	return nil
}

func (s *Handler) CHA(param int) error {
	*s.Events = append(*s.Events, CHA{param})
	return nil
}

func (s *Handler) VPA(param int) error {
	*s.Events = append(*s.Events, VPA{param})
	return nil
}

func (s *Handler) CUP(x int, y int) error {
	*s.Events = append(*s.Events, CUP{x, y})
	return nil
}

func (s *Handler) HVP(x int, y int) error {
	*s.Events = append(*s.Events, HVP{x, y})
	return nil
}

func (s *Handler) DECTCEM(visible bool) error {
	*s.Events = append(*s.Events, DECTCEM{visible})
	return nil
}

func (s *Handler) DECOM(visible bool) error {
	*s.Events = append(*s.Events, DECOM{visible})
	return nil
}

func (s *Handler) DECCOLM(use132 bool) error {
	*s.Events = append(*s.Events, DECCOLM{use132})
	return nil
}

func (s *Handler) ED(param int) error {
	*s.Events = append(*s.Events, ED{param})
	return nil
}

func (s *Handler) EL(param int) error {
	*s.Events = append(*s.Events, EL{param})
	return nil
}

func (s *Handler) IL(param int) error {
	*s.Events = append(*s.Events, IL{param})
	return nil
}

func (s *Handler) DL(param int) error {
	*s.Events = append(*s.Events, DL{param})
	return nil
}

func (s *Handler) ICH(param int) error {
	*s.Events = append(*s.Events, ICH{param})
	return nil
}

func (s *Handler) DCH(param int) error {
	*s.Events = append(*s.Events, DCH{param})
	return nil
}

func (s *Handler) SGR(params []int) error {
	*s.Events = append(*s.Events, SGR{params})
	return nil
}

func (s *Handler) SU(param int) error {
	*s.Events = append(*s.Events, SU{param})
	return nil
}

func (s *Handler) SD(param int) error {
	*s.Events = append(*s.Events, SD{param})
	return nil
}

func (s *Handler) DA(params []string) error {
	*s.Events = append(*s.Events, DA{params})
	return nil
}

func (s *Handler) DECSTBM(top int, bottom int) error {
	*s.Events = append(*s.Events, DECSTBM{top, bottom})
	return nil
}

func (s *Handler) RI() error {
	*s.Events = append(*s.Events, RI{})
	return nil
}

func (s *Handler) IND() error {
	*s.Events = append(*s.Events, IND{})
	return nil
}

func (s *Handler) Flush() error {
	*s.Events = append(*s.Events, Flush{})
	return nil
}
