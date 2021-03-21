package tx

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"server/utils"
	"strconv"
	"strings"
	"sync"
)

var lock *sync.Mutex = &sync.Mutex{}
var lock1 *sync.Mutex = &sync.Mutex{}
var config *utils.MyConfig

type Tx struct {
	TxId  string
	Topic string
	Msg   string
}

func IntTobytes(n int32) []byte {
	data := int32(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}
func BytesToInt(bts []byte) int32 {
	bytebuffer := bytes.NewBuffer(bts)
	var data int32
	binary.Read(bytebuffer, binary.BigEndian, &data)

	return data
}

func SaveTx(txdata *Tx) {
	lock.Lock()
	file, _ := os.OpenFile(filepath.Join(config.G_Tx_Dir, "current"), os.O_APPEND, 0755)
	defer file.Close()
	defer lock.Unlock()
	data, _ := json.Marshal(txdata)
	size := IntTobytes(int32(len(data)))
	file.Write(append(size, data...))
}
func GetTxDatasByTxId(txid string) ([]*Tx, error) {
	lock.Lock()
	file, _ := os.Open(filepath.Join(config.G_Tx_Dir, "finish"))
	defer file.Close()
	defer lock.Unlock()
	sizebuf := make([]byte, 4, 4)
	txs := make([]*Tx, 0, 0)
	record := false
	for {
		n, err := file.Read(sizebuf)
		if n == 0 || err != nil {
			return txs, nil
		}
		// abcdef
		if n != 4 {
			return txs, nil
		}
		data := make([]byte, BytesToInt(sizebuf))
		_tx := &Tx{}
		json.Unmarshal(data, _tx)
		if record {
			txs = append(txs, _tx)
			continue
		}
		if CompareTxId(txid, _tx.TxId) == 0 {
			record = true
		}
	}
}

var unfinishTxs []*Tx = make([]*Tx, 0, 0)

// 事务提交
func CommitTx(txid string) {
	lock1.Lock()
	for i := 0; i < len(unfinishTxs); i++ {
		if CompareTxId(unfinishTxs[i].TxId, txid) == 0 {
			SaveTx(unfinishTxs[i])
			if len(unfinishTxs) == 1 {
				unfinishTxs = make([]*Tx, 0, 0)
			} else {
				if len(unfinishTxs)-1 == i {
					unfinishTxs = unfinishTxs[:i]
				} else {
					unfinishTxs = append(unfinishTxs[:i], unfinishTxs[i+1:]...)
				}
			}
		}
	}
	defer lock1.Unlock()
}
func GetTxDataByTxId(txid string) (*Tx, error) {
	lock.Lock()
	file, _ := os.Open(filepath.Join(config.G_Tx_Dir, "finish"))
	defer file.Close()
	defer lock.Unlock()
	sizebuf := make([]byte, 4, 4)
	for {
		n, err := file.Read(sizebuf)
		if n == 0 || err != nil {
			return nil, errors.New("not found")
		}
		// abcdef
		if n != 4 {
			return nil, errors.New("not found")
		}
		data := make([]byte, BytesToInt(sizebuf))
		_tx := &Tx{}
		json.Unmarshal(data, _tx)
		if CompareTxId(txid, _tx.TxId) == 0 {
			return _tx, nil
		}
	}
}
func GetCurrentTxId() string {
	file, err := os.Open(filepath.Join(config.G_Tx_Dir, "current"))
	if err != nil {
		os.Exit(0)
	}
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	return string(line)
}
func CompareTxId(id1, id2 string) int8 {
	id1s := strings.Split(id1, "-")
	id2s := strings.Split(id2, "-")
	fmt.Println("=======", id1, id2)
	index1, _ := strconv.Atoi(id1s[0])
	index2, _ := strconv.Atoi(id2s[0])
	if index1 > index2 {
		return 1
	} else if index1 < index2 {
		return -1
	}

	code1, _ := strconv.Atoi(id1s[2])
	code2, _ := strconv.Atoi(id2s[2])
	if code1 > code2 {
		return 1
	} else {
		return -1
	}
	return 0
}
func Init(conf *utils.MyConfig) {
	config = conf
}
