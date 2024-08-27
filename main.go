package main

import (
	"a/TOsql"
	"a/fecthrating"
	"fmt"
	"log"
)

type stu struct {
	LC_ID     int
	USER_NAME string
	REAL_NAME string
	LC_RATING float32
}

func main() {
	// 确保数据库连接已经初始化
	if TOsql.DB == nil {
		log.Fatal("DB is nil")
	}

	// 执行查询
	rows, err := TOsql.DB.Query("SELECT LC_ID, USER_NAME, REAL_NAME, LC_RATING FROM LeedcodeRating")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var stus []stu
	for rows.Next() {
		var s stu
		err := rows.Scan(&s.LC_ID, &s.USER_NAME, &s.REAL_NAME, &s.LC_RATING)
		if err != nil {
			log.Fatal(err)
		}
		stus = append(stus, s)
	}

	// 检查行扫描是否有错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 打印结果
	for _, s := range stus {
		fmt.Printf("ID: %d, UserName: %s, RealName: %s, Rating: %.2f\n", s.LC_ID, s.USER_NAME, s.REAL_NAME, s.LC_RATING)
	}
	fmt.Println("Total:", len(stus))
	fmt.Println("开始更新成员分数.....")

	// 更新成员分数
	for _, s := range stus {
		newRating, err := fecthrating.FetchRating(s.USER_NAME)
		if err != nil {
			fmt.Println(s.REAL_NAME, "获取不到分数，原因：", err)
			continue
		}

		_, err = TOsql.DB.Exec("UPDATE LeedcodeRating SET LC_RATING = ? WHERE LC_ID = ?", newRating, s.LC_ID)
		if err != nil {
			log.Fatal(err)
			fmt.Println(s.REAL_NAME, "更新失败,updata异常")
		} else {
			fmt.Println("更新成功", s.REAL_NAME, "的分数更新为：", newRating, "变化幅度：", int(float32(newRating)-s.LC_RATING))
		}
	}
	TOsql.DB.Close()
}
