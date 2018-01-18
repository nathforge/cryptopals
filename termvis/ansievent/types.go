package ansievent

type Print struct{ B byte }
type Execute struct{ B byte }
type CUU struct{ Param int }
type CUD struct{ Param int }
type CUF struct{ Param int }
type CUB struct{ Param int }
type CNL struct{ Param int }
type CPL struct{ Param int }
type CHA struct{ Param int }
type VPA struct{ Param int }
type CUP struct{ X, Y int }
type HVP struct{ X, Y int }
type DECTCEM struct{ Visible bool }
type DECOM struct{ Visible bool }
type DECCOLM struct{ Use132 bool }
type ED struct{ Param int }
type EL struct{ Param int }
type IL struct{ Param int }
type DL struct{ Param int }
type ICH struct{ Param int }
type DCH struct{ Param int }
type SGR struct{ Params []int }
type SU struct{ Param int }
type SD struct{ Param int }
type DA struct{ Params []string }
type DECSTBM struct{ Top, Bottom int }
type RI struct{}
type IND struct{}
type Flush struct{}
