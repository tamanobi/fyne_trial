run:
	go run main.go

package: clean
	fyne package -icon ~/Desktop/icon.png

clean:
	rm -rf ./fyne_trial.app ./fyne_trial

.PHONY:	package clean run
