package tables

import (
	"database/sql"
	"log"
)

// CreateSeatTable 열람실 관리 테이블 및 인덱스를 생성합니다.
// 함수 이름을 대문자로 시작하여 외부에서 접근 가능하게 만듭니다.
func CreateSeatTable(db *sql.DB) error {
	log.Println("좌석 관리 테이블을 생성합니다...")

	// 테이블 생성
	createBaseTableQuery := `CREATE TABLE IF NOT EXISTS seat_table ();`

	_, err := db.Exec(createBaseTableQuery)
	if err != nil {
		return err
	}
	log.Println("seat_table 테이블 기본 구조 생성 완료")

	// 각 필드 개별 추가
	fieldQueries := []string{
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS code_number SERIAL PRIMARY KEY;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS reading_room_code INTEGER NOT NULL;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS seat_number INTEGER NOT NULL;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS power_number INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS display_power_number INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS password TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS member_id TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS member_name TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS memo TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS usage_month VARCHAR(6);`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS check_in_time TIME;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS check_in_button BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS cleaning_light BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS check_in_type INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS outing_datetime TIMESTAMP;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS seat_release_datetime TIMESTAMP;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS registration_date DATE;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS registration_time TIME;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS extension_datetime TIMESTAMP;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS expiration_date DATE;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS expiration_time TIME;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS m_top INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS m_left INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS m_width INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS m_height INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS card_number TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS remote_control_used BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS daily_remote_control_used INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS pb_height INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS grade_number INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS grade_name TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS another_name TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS gender CHAR(1);`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS unmanned_grade INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS unmanned_disabled BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS is_admin BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS f_top INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS f_left INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS f_width INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS f_height INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS free_seat BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS free_fixed_seat BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS free_waiting_seat BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS release_waiting_seat BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS free_seat_room BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS regular_fixed_seat BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS locker_used BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS exclude_cleaning BOOLEAN;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS r_top INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS r_left INTEGER;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS registration_type TEXT;`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS purchased_amount NUMERIC(10,2);`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS additional_amount NUMERIC(10,2);`,
		`ALTER TABLE seat_table ADD COLUMN IF NOT EXISTS move_grade INTEGER;`,
	}

	// 각 필드 추가 실행 및 진행 상황 로깅
	for i, query := range fieldQueries {
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
		// if (i+1)%10 == 0 || i == len(fieldQueries)-1 {
		log.Printf("seat_table 필드 추가 진행 중: %d/%d 완료", i+1, len(fieldQueries))
		// }
	}

	// 인덱스 생성 쿼리 목록 (이미 존재하면 생성하지 않음)
	indexQueries := []string{
		`CREATE INDEX IF NOT EXISTS idx_reading_room_seat ON seat_table (reading_room_code, seat_number);`,
		`CREATE INDEX IF NOT EXISTS idx_member_id ON seat_table (member_id);`,
		`CREATE INDEX IF NOT EXISTS idx_usage_month ON seat_table (usage_month);`,
		`CREATE INDEX IF NOT EXISTS idx_registration_date ON seat_table (registration_date);`,
		`CREATE INDEX IF NOT EXISTS idx_card_number ON seat_table (card_number);`,
		`CREATE INDEX IF NOT EXISTS idx_grade_number ON seat_table (grade_number);`,
		`CREATE INDEX IF NOT EXISTS idx_registration_type ON seat_table (registration_type);`,
		`CREATE INDEX IF NOT EXISTS idx_move_grade ON seat_table (move_grade);`,
		`CREATE INDEX IF NOT EXISTS idx_check_in_type ON seat_table (check_in_type);`,
		`CREATE INDEX IF NOT EXISTS idx_outing_datetime ON seat_table (outing_datetime);`,
		`CREATE INDEX IF NOT EXISTS idx_seat_release_datetime ON seat_table (seat_release_datetime);`,
	}

	// 인덱스 생성 실행
	for i, query := range indexQueries {
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
		log.Printf("seat_table 인덱스 추가 중: %d/%d 완료", i+1, len(indexQueries))
	}

	log.Println("seat_table 테이블과 인덱스가 성공적으로 생성되었습니다.")
	return nil
}
