package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"main/model"
	"main/view"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FortunesHandler struct{}

func (h FortunesHandler) HandleGetFortuneList(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	opened := c.QueryParam("opened") == "true"

	if opened {
		openedFortunes, err := model.GetOpened(database)
		if err != nil {
			return err
		}
		return render(c, view.FortuneOpenedList(openedFortunes))
	} else {
		memories, err := model.GetMemories(database)
		if err != nil {
			return err
		}
		wishes, err := model.GetWishes(database)
		if err != nil {
			return err
		}
		return render(c, view.FortuneList(memories, wishes))
	}
}

func (h FortunesHandler) HandleGetFortune(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	fortune, err := model.GetFortune(database, id)
	if err != nil {
		return err
	}
	if !fortune.Opened {
		_, err := model.OpenFortune(database, id)
		if err != nil {
			return err
		}
	}
	return render(c, view.Fortune(fortune))
}

func (h FortunesHandler) HandleNewFortunes(c echo.Context) error {
	kind := c.QueryParam("kind")
	if kind == string(model.KindMemory) {
		return render(c, view.FortunesForm(model.KindMemory))
	} else if kind == string(model.KindWish) {
		return render(c, view.FortunesForm(model.KindWish))
	}
	return fmt.Errorf("invalid kind")
}

func (h FortunesHandler) HandlePostFortunes(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	memory := c.FormValue("memory")
	wish := c.FormValue("wish")
	var err error
	if memory != "" {
		_, err = model.CreateFortune(database, model.KindMemory, memory)
	} else if wish != "" {
		_, err = model.CreateFortune(database, model.KindWish, wish)
	} else {
		return fmt.Errorf("no fortune content provided")
	}

	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func (h FortunesHandler) HandleOpenFortune(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	fortuneId, err := model.OpenFortune(database, id)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/fortunes/"+fmt.Sprintf("%d", fortuneId))
}

func (h FortunesHandler) HandleDeleteFortunes(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	err := model.DeleteAllFortunes(database)
	if err != nil {
		return errors.New("fortunes not deleted")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
