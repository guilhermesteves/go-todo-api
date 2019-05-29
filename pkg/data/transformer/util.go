package transformer

import (
	"github.com/guilhermesteves/go-todo-api/internal/pkg/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func toFormattedTime(d interface{}) string {
	switch d.(type) {
	case primitive.DateTime:
		t := time.Unix(int64(d.(primitive.DateTime))/1000, int64(d.(primitive.DateTime))%1000*1000000)
		return t.Format(util.DefaultTimeMask)
	case time.Time:
		return d.(time.Time).Format(util.DefaultTimeMask)
	}

	return ""
}
