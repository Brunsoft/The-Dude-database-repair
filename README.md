# The-Dude-database-repair
Procedura automatica per eseguire una pulizia dei database Dude (dude.db) a seguito dell'errore "database disk image is malformed"

> ## Istruzioni
  * Recuperare il proprio file `dude.db` dalla cartella `dude`.
  * Scaricare questa repository ed estrarne il contenuto.
  * Copiare all'interno il file `dude.db`.
  * Aprire il Prompt dei Comandi ed eseguire `dude.exe`
  * Al termine dell'operazione noterete che il file `dude.db` avrà una dimensione enormemente inferiore alla precedente. A questo punto sarà possibile utilizzare `dude.db` per ripristinare il database.
  * Dal terminale di Dude eseguire `/dude set enabled=no`
  * Sostituire quindi il proprio `dude.db` dalla cartella `dude` con quello generato da questa procedura.
  * Dal terminale di Dude eseguire `/dude set enabled=yes`
