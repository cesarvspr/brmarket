package brmarket

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type IndexType int

const (
	IBOV IndexType = iota
	IBXX
	IBXL
	IBRA
	IGCX
	ITAG
	IGNM
	IGCT
	IDIV
	MLCX
	IVBX
	ICO2
	ISEE
	ICON
	IEEX
	IFNC
	IMOB
	INDX
	IMAT
	UTIL
	IFIX
	BDRX
)

// Response type for B3 API
type B3IndexInfo struct {
	BizSts         bizSts           `json:"BizSts"`
	Index          index            `json:"Index"`
	UnderlyingList []underlyingList `json:"UnderlyingList"`
	Msg            msg              `json:"Msg"`
}

type bizSts struct {
	Cd string `json:"cd"`
}

type index struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

type underlyingList struct {
	IndexTheoreticalQty float64 `json:"indexTheoreticalQty"`
	IndxCmpnPctg        float64 `json:"indxCmpnPctg"`
	Symb                string  `json:"symb"`
	Desc                string  `json:"desc"`
}

type msg struct {
	DtTm string `json:"dtTm"`
}

// Errors
var (
	// ErrIndexNotFound is used for all indexes not supported
	ErrIndexNotFound = errors.New("index not found")
)

func (i IndexType) String() string {
	return [...]string{
		"IBOV",
		"IBXX",
		"IBXL",
		"IBRA",
		"IGCX",
		"ITAG",
		"IGNM",
		"IGCT",
		"IDIV",
		"MLCX",
		"IVBX",
		"ICO2",
		"ISEE",
		"ICON",
		"IEEX",
		"IFNC",
		"IMOB",
		"INDX",
		"IMAT",
		"UTIL",
		"IFIX",
		"BDRX",
	}[i]
}

func IndexFromString(input string) (*IndexType, error) {
	for i := IBOV; i <= BDRX; i++ {
		if i.String() == input {
			return &i, nil
		}
	}

	return nil, ErrIndexNotFound
}

// Return shares for reach index, only Brazilian indexes allowed.
// Example: GetSharesForIndex("IBOV") or GetSharesForIndex(brmarket.IBOV.String())
func GetSharesForIndex(index string) (B3IndexInfo, error) {

	idx, indexErr := IndexFromString(index)
	if indexErr != nil {
		return B3IndexInfo{}, indexErr
	}

	URL := "https://cotacao.b3.com.br/mds/api/v1/IndexComposition/" + idx.String()
	resp, err := http.Get(URL)
	if err != nil {
		return B3IndexInfo{}, err
	}
	defer resp.Body.Close()

	//Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return B3IndexInfo{}, err
	}
	var info B3IndexInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return info, err
	}

	return info, nil
}
