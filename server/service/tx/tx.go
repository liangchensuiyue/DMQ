package tx

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"server/utils"
	"strconv"
	"strings"
	"sync"
)

var lock *sync.Mutex = &sync.Mutex{}
var lock1 *sync.Mutex = &sync.Mutex{}
var lock2 *sync.Mutex = &sync.Mutex{}

// var lock2 *sync.Mutex = &sync.Mutex{}
var config *utils.MyConfig
var currentTxId string

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
	file, _ := os.OpenFile(filepath.Join(config.G_Tx_Dir, "finish_tx"), os.O_APPEND, 0755)
	defer file.Close()
	defer lock.Unlock()
	data, _ := json.Marshal(txdata)
	size := IntTobytes(int32(len(data)))
	file.Write(append(size, data...))
}
func GetTxDatasByTxId(txid string) ([]*Tx, error) {
	lock.Lock()
	file, _ := os.Open(filepath.Join(config.G_Tx_Dir, "finish_tx"))
	defer lock.Unlock()
	sizebuf := make([]byte, 4, 4)
	txs := make([]*Tx, 0, 0)
	record := false
	if txid == "" {
		record = true
	}
	for {
		n, err := file.Read(sizebuf)
		if n == 0 || err != nil {
			file.Close()
			return txs, nil
		}
		// abcdef
		if n != 4 {
			file.Close()
			return txs, nil
		}
		data := make([]byte, BytesToInt(sizebuf))
		file.Read(data)
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

func PrepareTx(data *Tx) {
	lock1.Lock()
	defer lock1.Unlock()
	for _, v := range unfinishTxs {
		if v.TxId == data.TxId {
			return
		}
	}
	unfinishTxs = append(unfinishTxs, data)
}

// 事务提交
func CommitTx(txid string) error {
	lock1.Lock()
	defer lock1.Unlock()

	for i := 0; i < len(unfinishTxs); i++ {
		if CompareTxId(unfinishTxs[i].TxId, txid) == 0 {
			if len(unfinishTxs) == 1 {
				unfinishTxs = make([]*Tx, 0, 0)
			} else {
				if len(unfinishTxs)-1 == i {
					unfinishTxs = unfinishTxs[:i]
				} else {
					unfinishTxs = append(unfinishTxs[:i], unfinishTxs[i+1:]...)
				}
			}
			return nil
		}
	}
	return errors.New("not found")
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

	return currentTxId
}
func GetCurrentTxIdFromDisk() {
	lock2.Lock()
	defer lock2.Unlock()
	file, err := os.Open(filepath.Join(config.G_Tx_Dir, "current"))

	if err != nil {
		os.Exit(0)
	}
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	currentTxId = strings.TrimSpace(string(line))
	file.Close()
}
func WriteCurrentTxId(txid string) {
	lock2.Lock()
	defer lock2.Unlock()
	file, err := os.OpenFile(filepath.Join(config.G_Tx_Dir, "current"), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		os.Exit(0)
	}
	file.Write([]byte(txid))
	currentTxId = txid
	file.Close()
}
func CompareTxId(id1, id2 string) int8 {
	if id1 == "" && id2 == "" {
		return 0
	}
	if id1 == "" && id2 != "" {
		return -1
	}
	if id1 != "" && id2 == "" {
		return 1
	}
	id1s := strings.Split(id1, "-")
	id2s := strings.Split(id2, "-")
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
	} else if code1 < code2 {
		return -1
	}
	return 0
}
func Init(conf *utils.MyConfig) {
	config = conf
	GetCurrentTxIdFromDisk()
}
