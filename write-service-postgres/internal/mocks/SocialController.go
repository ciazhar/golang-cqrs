// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gin "github.com/gin-gonic/gin"
import mock "github.com/stretchr/testify/mock"

// SocialController is an autogenerated mock type for the SocialController type
type SocialController struct {
	mock.Mock
}

// Delete provides a mock function with given fields: c
func (_m *SocialController) Delete(c *gin.Context) {
	_m.Called(c)
}

// Store provides a mock function with given fields: c
func (_m *SocialController) Store(c *gin.Context) {
	_m.Called(c)
}

// Update provides a mock function with given fields: c
func (_m *SocialController) Update(c *gin.Context) {
	_m.Called(c)
}