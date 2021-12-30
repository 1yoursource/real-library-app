package auto_incrementer

import "lib-client-server/database"

var AI = create(database.Connect("localhost","libDB", "aiRecording").C("ai"))