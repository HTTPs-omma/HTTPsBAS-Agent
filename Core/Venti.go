package Core

import (
	"os"
	"strconv"
	"strings"
)

type Venti struct {
}

func (vt *Venti) ReplaceStringOne(s1 string, files []string, update string) string {
	s1 = vt.replaceAgentSig(s1)
	s1 = vt.replaceFileSig(s1, files)
	s1 = vt.replaceUpdateSig(s1, update)

	return s1
}

func (vt *Venti) ReplaceString(s1 string, s2 string, files []string, update string) (string, string) { // 명명된 반환값
	s1 = vt.replaceAgentSig(s1)
	s2 = vt.replaceAgentSig(s2)
	s1 = vt.replaceFileSig(s1, files)
	s2 = vt.replaceFileSig(s2, files)
	s1 = vt.replaceUpdateSig(s1, update)
	s2 = vt.replaceUpdateSig(s2, update)

	return s1, s2
}

func (vt *Venti) replaceAgentSig(str string) string {
	return strings.Replace(str, "@agent.exe", os.Getenv("exePath"), -1)
}

func (vt *Venti) replaceFileSig(str string, files []string) string {
	for i, v := range files {
		i += 1
		str = strings.Replace(str, "@file"+strconv.Itoa(i), v, -1)
	}
	return str
}

func (vt *Venti) replaceUpdateSig(str string, update string) string {
	str = strings.Replace(str, "@upload", update, -1)
	return str
}
