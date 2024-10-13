LOCAL_PATH := $(call my-dir)
include $(CLEAR_VARS)

LOCAL_MODULE    := avoengine
LOCAL_SRC_FILES := $(call all-subdir-cpp-files)
LOCAL_LDLIBS    := -llog -landroid -lEGL -lGLESv2

include $(BUILD_SHARED_LIBRARY)
