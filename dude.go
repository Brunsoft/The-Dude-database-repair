package main

import(
	"io"
    "fmt"
    "os"
    "os/exec"
    "io/ioutil"
    "bufio"
    "strings"
)

func execCommand(command string) {
	c := exec.Command("cmd", "/C", command)

    if err := c.Run(); err != nil { 
        fmt.Println("Error: ", err)
    }
    fmt.Println("   > comando eseguito con successo") 
}


func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if err != nil { return }

	fmt.Println("   > cancellazione file eseguita")
}

func copyFile(src, dst string) {
        sourceFileStat, err := os.Stat(src)
        if err != nil { return }

        if !sourceFileStat.Mode().IsRegular() { return }

        source, err := os.Open(src)
        if err != nil { return }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil { return }
        
        defer destination.Close()
        _, err = io.Copy(destination, source)
        if err != nil { return }
        
        fmt.Println("   > copia file eseguita")
        return 
}

func main(){ 
	fmt.Println(" 0: Apro dude.db con Sqlite3 ed eseguo il vacuum ... ")  
	execCommand("echo vacuum; | sqlite3.exe dude.db");
 
    fmt.Println(" 1: Creo il file dude.sql a partire da dude.db ... ") 
    execCommand("echo .dump | sqlite3.exe dude.db > dude.sql");  

	fmt.Println(" 2: Apro il file dude.sql ... ") 
    f, err := os.Open("dude.sql")
	if err != nil { return }
	defer f.Close()
	fmt.Println("   > dude.sql aperto con successo")

	fmt.Println(" 3: Creo il file objs.sql ... ")
	err = ioutil.WriteFile("objs.sql", []byte(""), 0777)
    if err != nil { return }
    fmt.Println("   > objs.sql creato con successo")
    
    fmt.Println(" 4: Apro il file objs.sql ... ") 
    f1, err := os.OpenFile("objs.sql", os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil { return }
    defer f1.Close()
    fmt.Println("   > objs.sql aperto con successo")
    
	scanner := bufio.NewScanner(f)
	line := 0
	fmt.Println(" 5: Analizzo dude.sql ... ") 
	for scanner.Scan() {
		row := scanner.Text()
	    if strings.Contains(row, "INSERT INTO objs VALUES") {
	    	line++
	    	f1.WriteString(row + "\n");
	    }
	}
	fmt.Println("   >", line, " righe copiate in objs.sql")
	
    fmt.Println(" 6: Cancello l'originale file dude.db ... ")
    deleteFile("dude.db");
    
    fmt.Println(" 7: Copio un file dude.db vergine ... ")
    copyFile("files/dude.db", "dude.db");
	
	fmt.Println(" 8: Elimino le righe objs dal nuovo dude.db ... ")  
	execCommand("echo delete from objs; | sqlite3.exe dude.db");
	
	fmt.Println(" 9: Popolo il nuovo db con le vecchie righe objs ... ")   
	execCommand("echo .read objs.sql | sqlite3.exe dude.db");
	
	fmt.Println("\n Procedura completata!");
    
}