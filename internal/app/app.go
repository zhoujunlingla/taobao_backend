package app

import (
	"fmt"
	"runtime"
	"taobao_backend/config"
	"taobao_backend/internal/database"
	"taobao_backend/internal/routes"
)

func Run() {
	config.Init()
	printHello()
	database.InitMysql()
	routes.InitRouter()
}

func printHello() {
	fmt.Println("#-----------------------------------------------------#")
	fmt.Println("# start taobao_backend with", runtime.Version())
	fmt.Println("#")
	fmt.Println("# *****  *****  *        *    *   *  ***** ")
	fmt.Println("# *      *   *  *       * *   **  *  *     ")
	fmt.Println("# *  **  *   *  *      *****  * * *  *  ** ")
	fmt.Println("# *   *  *   *  *      *   *  *  **  *   * ")
	fmt.Println("# *****  *****  *****  *   *  *   *  ***** ")
	fmt.Println("#")
	fmt.Println("#-----------------------------------------------------#")
	fmt.Println()
}
