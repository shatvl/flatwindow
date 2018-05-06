package seed

import (
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/repositories"
)

func SeedAgents() {
	session := mongo.Session()
	defer session.Close()

	tsAgent := &models.User{
		Email: "mik@t-s.by",
		Role: config.ROLE_AGENT,
		Password: "tew2tqQ_1",
		AgentType: repositories.TsType,
		AgentCode: repositories.FeedTypeToName[repositories.TsType],
	}
	_, _ = repositories.NewUserRepository().Create(tsAgent)
}