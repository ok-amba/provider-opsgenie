# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Targets

img.publish:
	@$(INFO) Skipping image publish for $(IMAGE)
	@echo Publish is deferred to xpkg machinery
	@$(OK) Image publish skipped for $(IMAGE)

img.build:
	@$(INFO) docker build $(IMAGE)
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp -R ../../../package $(IMAGE_TEMP_DIR) || $(FAIL)
	@cd $(IMAGE_TEMP_DIR) && find package -type f -name '*.yaml' -exec cat {} >> 'package.yaml' \; -exec printf '\n---\n' \; || $(FAIL)
	@docker build -t $(IMAGE) $(IMAGE_TEMP_DIR) || $(FAIL)
	@docker tag $(IMAGE) "$(IMAGE):$(VERSION)" || $(FAIL)
	@$(OK) docker build $(IMAGE)
