package tables

import (
	"database/sql"
	"log"
)

// CreateRoomTable 좌석 관리 테이블 및 인덱스를 생성합니다.
// 함수 이름을 대문자로 시작하여 외부에서 접근 가능하게 만듭니다.
func CreateRoomTable(db *sql.DB) error {
	log.Println("열람실 관리 테이블을 생성합니다...")

	// 테이블 생성
	createBaseTableQuery := `CREATE TABLE IF NOT EXISTS room_table ();`

	_, err := db.Exec(createBaseTableQuery)
	if err != nil {
		return err
	}
	log.Println("room_table 테이블 기본 구조 생성 완료")

	// 각 필드 개별 추가
	fieldQueries := []string{
		//테이블
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS auto_increment SERIAL PRIMARY KEY;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS chain_code SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS comapny_code SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS room_code SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS room_title VARCHAR(20) NOT NULL;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS title_background_color VARCHAR(9);`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS title_text_color VARCHAR(9);`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS room_background_color VARCHAR(9);`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS room_top INTEGER;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS room_left INTEGER;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS room_width INTEGER;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS room_height INTEGER;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS gender SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS waiting SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS release SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS hide_title SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS transparent_background SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS hide_border SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS kiosk_disabled SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS power_control SMALLINT;`,
		`ALTER TABLE room_table ADD COLUMN IF NOT EXISTS breaker_number INTEGER;`,
	}

	// 각 필드 추가 실행 및 진행 상황 로깅
	for i, query := range fieldQueries {
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
		// if (i+1)%5 == 0 || i == len(fieldQueries)-1 {
		log.Printf("room_table 필드 추가 진행 중: %d/%d 완료", i+1, len(fieldQueries))
		// }
	}

	// 인덱스 생성 쿼리 목록
	indexQueries := []string{
		`CREATE INDEX IF NOT EXISTS idx_chain_code ON room_table (chain_code);`,
		`CREATE INDEX IF NOT EXISTS idx_comapny_code ON room_table (comapny_code);`,
		`CREATE INDEX IF NOT EXISTS idx_room_code ON room_table (room_code);`,
	}

	// 인덱스 생성 실행
	for _, query := range indexQueries {
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}

	log.Println("room_table 테이블과 인덱스가 성공적으로 생성되었습니다.")
	return nil
}
