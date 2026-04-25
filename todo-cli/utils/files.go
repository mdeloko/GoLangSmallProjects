package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/mdeloko/GoLangSmallProjects/todo-cli/model"
)

func GetJSONDatabase()(file *os.File,err error){
	path := "./json/db.json"

	dir := filepath.Dir(path)
	if err = os.MkdirAll(dir,0755) ; err != nil {
		return nil, fmt.Errorf("Falha ao criar diretório %s - %w",dir,err)
	}

	file, err = os.OpenFile(path,os.O_RDWR|os.O_CREATE,0644)
	if err != nil {
        return nil, fmt.Errorf("Falha ao abrir ou criar o arquivo json: %w", err)
    }

	return file, nil
}

func UpdateTasks(file *os.File, tasks []model.Task) (err error) {
	if _,err = file.Seek(0,0); err != nil {
		return fmt.Errorf("Falha ao mover o cursor: %w",err)
	}
	if err := file.Truncate(0); err != nil {
        return fmt.Errorf("Falha ao limpar o arquivo original: %w", err)
    }
	encoder := json.NewEncoder(file)
	encoder.SetIndent("","    ")
	if err := encoder.Encode(tasks); err != nil {
        return fmt.Errorf("falha ao salvar struct Go para JSON: %w", err)
    }
	return nil
}

func ReadTasks(file *os.File,taskList *[]model.Task)(err error){
	if _,err = file.Seek(0,0); err != nil {
		return fmt.Errorf("Falha ao mover o cursor: %w",err)
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		return nil
	}
	
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&taskList); err != nil {
		return fmt.Errorf("Falha ao transformar JSON em Struct Go: %w",err)
	}
	return nil
}