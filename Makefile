# To test, build, deploy offline-tasks
#   -: ignore this commnad error
#   @: no display current commnad to std output

# Commnads declare
GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

# Params define
MAIN_PATH=../main
PACKAGE_PATH=../package
PACKAGE_BIN_PATH=../package/bin
BIN=offline-tasks
FILENAME=offline-tasks.tar.gz

# Deploy Params
DEV_HOST=zy-dev
DEV_TAR_PATH=/home/worker/project/offline-tasks

PROD_HOST=zy-pro2
PROD_TAR_PATH=/home/worker/project/offline-tasks

default: build pack

test:
  # testing
  - $(GOTEST) ../... -v

build:
  # building
  mkdir $(PACKAGE_PATH)
  mkdir $(PACKAGE_BIN_PATH)
  cd $(MAIN_PATH) && $(GOBUILD) -o $(BIN)
  mv "$(MAIN_PATH)/$(BIN)" $(PACKAGE_BIN_PATH)
  cp -r "../configs" $(PACKAGE_PATH)
  cp "../sh/start.sh" $(PACKAGE_BIN_PATH)

pack:
  # packing
  cd $(PACKAGE_PATH) && tar -zcvf ../$(FILENAME) ./*
  mv ../$(FILENAME) $(PACKAGE_PATH)


##################################################
#                                                #
#   deploy:         from zy-dev to execute       #
#   deploy-dev:     from dev-CI to execute       #
#   deploy-prod:    from prod-CI to execute      #
#                                                #
##################################################

deploy: clean build pack
  # deploy dev from dev
  cp $(PACKAGE_PATH)/$(FILENAME) $(DEV_TAR_PATH)
  cd $(DEV_TAR_PATH) && tar zxvf $(FILENAME) && supervisorctl -c configs/dev.supervisord.conf restart offline-tasks

deploy-dev: clean build pack
  # deploy-dev from CI
  scp $(PACKAGE_PATH)/$(FILENAME) $(DEV_HOST):$(DEV_TAR_PATH)
  ssh $(DEV_HOST) "cd $(DEV_TAR_PATH) && tar zxvf $(FILENAME) && supervisorctl -c configs/dev.supervisord.conf restart offline-tasks"

deploy-prod: clean build pack
  # deploying prod from dev or CI
  scp $(PACKAGE_PATH)/$(FILENAME) $(PROD_HOST):$(PROD_TAR_PATH)
  ssh $(PROD_HOST) "cd $(PROD_TAR_PATH) && tar zxvf $(FILENAME) && supervisorctl -c configs/prod.supervisord.conf restart offline-tasks"

clean:
  # cleaning
  rm -fr $(PACKAGE_PATH)
  rm -fr ../$(FILENAME)
