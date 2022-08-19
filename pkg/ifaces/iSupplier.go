package ifaces

import (
	"github.com/allanpk716/ChineseSubFinder/pkg/settings"
	"github.com/allanpk716/ChineseSubFinder/pkg/types/series"
	"github.com/allanpk716/ChineseSubFinder/pkg/types/supplier"
	"github.com/sirupsen/logrus"
)

type ISupplier interface {
	CheckAlive(proxySettings ...*settings.ProxySettings) (bool, int64)

	IsAlive() bool

	GetSupplierName() string

	OverDailyDownloadLimit() bool

	GetLogger() *logrus.Logger

	GetSettings() *settings.Settings

	GetSubListFromFile4Movie(filePath string) ([]supplier.SubInfo, error)

	GetSubListFromFile4Series(seriesInfo *series.SeriesInfo) ([]supplier.SubInfo, error)

	GetSubListFromFile4Anime(seriesInfo *series.SeriesInfo) ([]supplier.SubInfo, error)
}
