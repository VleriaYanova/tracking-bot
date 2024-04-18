package handlers

import (
	"fmt"
	"log"
)

var twoYearsSellLink = "https://fr.mos.ru/pokupka-nedvizhimosti-dlya-vseh/ajax.php?category[]=PARTICIPANTS&y2_sell=1&status[]=PROCESSING&status[]=FINISHED&price_min=0&price_max=100000000000000&price_m_min=0&price_m_max=100000000000&area_min=0&area_max=100000000&floor_min=-1&floor_max=100000000&open_sale=0&pagesize=100000"
var inMomentSellLink = "https://fr.mos.ru/pokupka-nedvizhimosti-dlya-vseh/ajax.php?category[]=PARTICIPANTS&for_sell=1&status[]=PROCESSING&status[]=FINISHED&price_min=0&price_max=100000000000000&price_m_min=0&price_m_max=100000000000&area_min=0&area_max=100000000&floor_min=-1&floor_max=100000000&open_sale=0&pagesize=100000"

func (h *TrackingHandler) ApartmentsHandler(link string, eventName string) {
	log.Println("Processing...")
	// Get all appartments from site

	outerApps, err := h.appService.GetApartments(link)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Saves appartment in database if it not exists
	// and if appartment was created - notify all bot subscribers
	for _, outApp := range *outerApps {
		eventType := h.GetEventTypeByApp(&outApp)
		created, err := h.appService.CreateIfNotExist(&outApp, eventType)
		if err != nil {
			fmt.Println(err.Error())
		}
		if created {
			h.NotifyAllSubscribers(&outApp, AppAdded, eventType)
			continue
		}
		inApp, err := h.appService.GetById(outApp.ID, eventType)
		if err != nil {
			fmt.Println(err.Error())
		}
		if inApp.Requested == outApp.Requested {
			continue
		}
		_, err = h.appService.Update(&outApp, inApp.ID, eventType)
		if err != nil {
			fmt.Println(err.Error())
		}
		if outApp.Requested == 1 {
			h.NotifyAllSubscribers(&outApp, AppStatusFirstDeclaration, eventType)
		} else if inApp.Requested < 2 && outApp.Requested > 1 {
			h.NotifyAllSubscribers(&outApp, AppStatusAuction, eventType)
		}
	}

	// Remove appartment from database if it was deleted from site
	// and notify all bot subscribers
	removedApps, err := h.appService.RemoveDeletedApps(outerApps, eventName)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, rmApp := range *removedApps {
		eventType := h.GetEventTypeByApp(&rmApp)
		h.NotifyAllSubscribers(&rmApp, AppRemoved, eventType)
	}
}
