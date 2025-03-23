package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// 고유 ID
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		// 사용자 이름
		field.String("name").Default("unknown"),
		// 이메일 주소 (유니크)
		field.String("email").Unique(),
		// 비밀번호 (민감 정보)
		field.String("password").Sensitive(),
		field.String("salt").Sensitive(),
		// 역할 (관리자, 개발자, 엔지니어 등)
		field.Enum("role").Values("admin", "developer", "engineer").Default("developer"),
		// 활성화 상태
		field.Bool("active").Default(true),
		// 생성 및 업데이트 타임스탬프
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		// 마지막 로그인 시간
		field.Time("last_login").Nillable().Optional(),
		// 프로필 이미지 URL
		field.String("profile_image").Nillable().Optional(),
		// GitHub 토큰 해시 (민감 정보)
		field.String("github_token_hash").Sensitive().Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// GeneralSpec 권한 관리 (중간 테이블과의 관계)
		edge.To("permissions", UserGeneralSpecPermissions.Type),
		// Project와의 관계 추가
		edge.To("projects", Project.Type),
	}
}
