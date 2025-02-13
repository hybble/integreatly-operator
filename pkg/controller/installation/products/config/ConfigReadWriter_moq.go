// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package config

import (
	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"sync"
)

var (
	lockConfigReadWriterMockGetOperatorNamespace sync.RWMutex
	lockConfigReadWriterMockReadAMQOnline        sync.RWMutex
	lockConfigReadWriterMockReadAMQStreams       sync.RWMutex
	lockConfigReadWriterMockReadCodeReady        sync.RWMutex
	lockConfigReadWriterMockReadFuse             sync.RWMutex
	lockConfigReadWriterMockReadLauncher         sync.RWMutex
	lockConfigReadWriterMockReadNexus            sync.RWMutex
	lockConfigReadWriterMockReadProduct          sync.RWMutex
	lockConfigReadWriterMockReadRHSSO            sync.RWMutex
	lockConfigReadWriterMockReadSolutionExplorer sync.RWMutex
	lockConfigReadWriterMockReadThreeScale       sync.RWMutex
	lockConfigReadWriterMockWriteConfig          sync.RWMutex
	lockConfigReadWriterMockreadConfigForProduct sync.RWMutex
)

// Ensure, that ConfigReadWriterMock does implement ConfigReadWriter.
// If this is not the case, regenerate this file with moq.
var _ ConfigReadWriter = &ConfigReadWriterMock{}

// ConfigReadWriterMock is a mock implementation of ConfigReadWriter.
//
//     func TestSomethingThatUsesConfigReadWriter(t *testing.T) {
//
//         // make and configure a mocked ConfigReadWriter
//         mockedConfigReadWriter := &ConfigReadWriterMock{
//             GetOperatorNamespaceFunc: func() string {
// 	               panic("mock out the GetOperatorNamespace method")
//             },
//             ReadAMQOnlineFunc: func() (*AMQOnline, error) {
// 	               panic("mock out the ReadAMQOnline method")
//             },
//             ReadAMQStreamsFunc: func() (*AMQStreams, error) {
// 	               panic("mock out the ReadAMQStreams method")
//             },
//             ReadCodeReadyFunc: func() (*CodeReady, error) {
// 	               panic("mock out the ReadCodeReady method")
//             },
//             ReadFuseFunc: func() (*Fuse, error) {
// 	               panic("mock out the ReadFuse method")
//             },
//             ReadLauncherFunc: func() (*Launcher, error) {
// 	               panic("mock out the ReadLauncher method")
//             },
//             ReadNexusFunc: func() (*Nexus, error) {
// 	               panic("mock out the ReadNexus method")
//             },
//             ReadProductFunc: func(product v1alpha1.ProductName) (ConfigReadable, error) {
// 	               panic("mock out the ReadProduct method")
//             },
//             ReadRHSSOFunc: func() (*RHSSO, error) {
// 	               panic("mock out the ReadRHSSO method")
//             },
//             ReadSolutionExplorerFunc: func() (*SolutionExplorer, error) {
// 	               panic("mock out the ReadSolutionExplorer method")
//             },
//             ReadThreeScaleFunc: func() (*ThreeScale, error) {
// 	               panic("mock out the ReadThreeScale method")
//             },
//             WriteConfigFunc: func(config ConfigReadable) error {
// 	               panic("mock out the WriteConfig method")
//             },
//             readConfigForProductFunc: func(product v1alpha1.ProductName) (ProductConfig, error) {
// 	               panic("mock out the readConfigForProduct method")
//             },
//         }
//
//         // use mockedConfigReadWriter in code that requires ConfigReadWriter
//         // and then make assertions.
//
//     }
type ConfigReadWriterMock struct {
	// GetOperatorNamespaceFunc mocks the GetOperatorNamespace method.
	GetOperatorNamespaceFunc func() string

	// ReadAMQOnlineFunc mocks the ReadAMQOnline method.
	ReadAMQOnlineFunc func() (*AMQOnline, error)

	// ReadAMQStreamsFunc mocks the ReadAMQStreams method.
	ReadAMQStreamsFunc func() (*AMQStreams, error)

	// ReadCodeReadyFunc mocks the ReadCodeReady method.
	ReadCodeReadyFunc func() (*CodeReady, error)

	// ReadFuseFunc mocks the ReadFuse method.
	ReadFuseFunc func() (*Fuse, error)

	// ReadLauncherFunc mocks the ReadLauncher method.
	ReadLauncherFunc func() (*Launcher, error)

	// ReadNexusFunc mocks the ReadNexus method.
	ReadNexusFunc func() (*Nexus, error)

	// ReadProductFunc mocks the ReadProduct method.
	ReadProductFunc func(product v1alpha1.ProductName) (ConfigReadable, error)

	// ReadRHSSOFunc mocks the ReadRHSSO method.
	ReadRHSSOFunc func() (*RHSSO, error)

	// ReadSolutionExplorerFunc mocks the ReadSolutionExplorer method.
	ReadSolutionExplorerFunc func() (*SolutionExplorer, error)

	// ReadThreeScaleFunc mocks the ReadThreeScale method.
	ReadThreeScaleFunc func() (*ThreeScale, error)

	// WriteConfigFunc mocks the WriteConfig method.
	WriteConfigFunc func(config ConfigReadable) error

	// readConfigForProductFunc mocks the readConfigForProduct method.
	readConfigForProductFunc func(product v1alpha1.ProductName) (ProductConfig, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetOperatorNamespace holds details about calls to the GetOperatorNamespace method.
		GetOperatorNamespace []struct {
		}
		// ReadAMQOnline holds details about calls to the ReadAMQOnline method.
		ReadAMQOnline []struct {
		}
		// ReadAMQStreams holds details about calls to the ReadAMQStreams method.
		ReadAMQStreams []struct {
		}
		// ReadCodeReady holds details about calls to the ReadCodeReady method.
		ReadCodeReady []struct {
		}
		// ReadFuse holds details about calls to the ReadFuse method.
		ReadFuse []struct {
		}
		// ReadLauncher holds details about calls to the ReadLauncher method.
		ReadLauncher []struct {
		}
		// ReadNexus holds details about calls to the ReadNexus method.
		ReadNexus []struct {
		}
		// ReadProduct holds details about calls to the ReadProduct method.
		ReadProduct []struct {
			// Product is the product argument value.
			Product v1alpha1.ProductName
		}
		// ReadRHSSO holds details about calls to the ReadRHSSO method.
		ReadRHSSO []struct {
		}
		// ReadSolutionExplorer holds details about calls to the ReadSolutionExplorer method.
		ReadSolutionExplorer []struct {
		}
		// ReadThreeScale holds details about calls to the ReadThreeScale method.
		ReadThreeScale []struct {
		}
		// WriteConfig holds details about calls to the WriteConfig method.
		WriteConfig []struct {
			// Config is the config argument value.
			Config ConfigReadable
		}
		// readConfigForProduct holds details about calls to the readConfigForProduct method.
		readConfigForProduct []struct {
			// Product is the product argument value.
			Product v1alpha1.ProductName
		}
	}
}

// GetOperatorNamespace calls GetOperatorNamespaceFunc.
func (mock *ConfigReadWriterMock) GetOperatorNamespace() string {
	if mock.GetOperatorNamespaceFunc == nil {
		panic("ConfigReadWriterMock.GetOperatorNamespaceFunc: method is nil but ConfigReadWriter.GetOperatorNamespace was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockGetOperatorNamespace.Lock()
	mock.calls.GetOperatorNamespace = append(mock.calls.GetOperatorNamespace, callInfo)
	lockConfigReadWriterMockGetOperatorNamespace.Unlock()
	return mock.GetOperatorNamespaceFunc()
}

// GetOperatorNamespaceCalls gets all the calls that were made to GetOperatorNamespace.
// Check the length with:
//     len(mockedConfigReadWriter.GetOperatorNamespaceCalls())
func (mock *ConfigReadWriterMock) GetOperatorNamespaceCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockGetOperatorNamespace.RLock()
	calls = mock.calls.GetOperatorNamespace
	lockConfigReadWriterMockGetOperatorNamespace.RUnlock()
	return calls
}

// ReadAMQOnline calls ReadAMQOnlineFunc.
func (mock *ConfigReadWriterMock) ReadAMQOnline() (*AMQOnline, error) {
	if mock.ReadAMQOnlineFunc == nil {
		panic("ConfigReadWriterMock.ReadAMQOnlineFunc: method is nil but ConfigReadWriter.ReadAMQOnline was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadAMQOnline.Lock()
	mock.calls.ReadAMQOnline = append(mock.calls.ReadAMQOnline, callInfo)
	lockConfigReadWriterMockReadAMQOnline.Unlock()
	return mock.ReadAMQOnlineFunc()
}

// ReadAMQOnlineCalls gets all the calls that were made to ReadAMQOnline.
// Check the length with:
//     len(mockedConfigReadWriter.ReadAMQOnlineCalls())
func (mock *ConfigReadWriterMock) ReadAMQOnlineCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadAMQOnline.RLock()
	calls = mock.calls.ReadAMQOnline
	lockConfigReadWriterMockReadAMQOnline.RUnlock()
	return calls
}

// ReadAMQStreams calls ReadAMQStreamsFunc.
func (mock *ConfigReadWriterMock) ReadAMQStreams() (*AMQStreams, error) {
	if mock.ReadAMQStreamsFunc == nil {
		panic("ConfigReadWriterMock.ReadAMQStreamsFunc: method is nil but ConfigReadWriter.ReadAMQStreams was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadAMQStreams.Lock()
	mock.calls.ReadAMQStreams = append(mock.calls.ReadAMQStreams, callInfo)
	lockConfigReadWriterMockReadAMQStreams.Unlock()
	return mock.ReadAMQStreamsFunc()
}

// ReadAMQStreamsCalls gets all the calls that were made to ReadAMQStreams.
// Check the length with:
//     len(mockedConfigReadWriter.ReadAMQStreamsCalls())
func (mock *ConfigReadWriterMock) ReadAMQStreamsCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadAMQStreams.RLock()
	calls = mock.calls.ReadAMQStreams
	lockConfigReadWriterMockReadAMQStreams.RUnlock()
	return calls
}

// ReadCodeReady calls ReadCodeReadyFunc.
func (mock *ConfigReadWriterMock) ReadCodeReady() (*CodeReady, error) {
	if mock.ReadCodeReadyFunc == nil {
		panic("ConfigReadWriterMock.ReadCodeReadyFunc: method is nil but ConfigReadWriter.ReadCodeReady was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadCodeReady.Lock()
	mock.calls.ReadCodeReady = append(mock.calls.ReadCodeReady, callInfo)
	lockConfigReadWriterMockReadCodeReady.Unlock()
	return mock.ReadCodeReadyFunc()
}

// ReadCodeReadyCalls gets all the calls that were made to ReadCodeReady.
// Check the length with:
//     len(mockedConfigReadWriter.ReadCodeReadyCalls())
func (mock *ConfigReadWriterMock) ReadCodeReadyCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadCodeReady.RLock()
	calls = mock.calls.ReadCodeReady
	lockConfigReadWriterMockReadCodeReady.RUnlock()
	return calls
}

// ReadFuse calls ReadFuseFunc.
func (mock *ConfigReadWriterMock) ReadFuse() (*Fuse, error) {
	if mock.ReadFuseFunc == nil {
		panic("ConfigReadWriterMock.ReadFuseFunc: method is nil but ConfigReadWriter.ReadFuse was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadFuse.Lock()
	mock.calls.ReadFuse = append(mock.calls.ReadFuse, callInfo)
	lockConfigReadWriterMockReadFuse.Unlock()
	return mock.ReadFuseFunc()
}

// ReadFuseCalls gets all the calls that were made to ReadFuse.
// Check the length with:
//     len(mockedConfigReadWriter.ReadFuseCalls())
func (mock *ConfigReadWriterMock) ReadFuseCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadFuse.RLock()
	calls = mock.calls.ReadFuse
	lockConfigReadWriterMockReadFuse.RUnlock()
	return calls
}

// ReadLauncher calls ReadLauncherFunc.
func (mock *ConfigReadWriterMock) ReadLauncher() (*Launcher, error) {
	if mock.ReadLauncherFunc == nil {
		panic("ConfigReadWriterMock.ReadLauncherFunc: method is nil but ConfigReadWriter.ReadLauncher was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadLauncher.Lock()
	mock.calls.ReadLauncher = append(mock.calls.ReadLauncher, callInfo)
	lockConfigReadWriterMockReadLauncher.Unlock()
	return mock.ReadLauncherFunc()
}

// ReadLauncherCalls gets all the calls that were made to ReadLauncher.
// Check the length with:
//     len(mockedConfigReadWriter.ReadLauncherCalls())
func (mock *ConfigReadWriterMock) ReadLauncherCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadLauncher.RLock()
	calls = mock.calls.ReadLauncher
	lockConfigReadWriterMockReadLauncher.RUnlock()
	return calls
}

// ReadNexus calls ReadNexusFunc.
func (mock *ConfigReadWriterMock) ReadNexus() (*Nexus, error) {
	if mock.ReadNexusFunc == nil {
		panic("ConfigReadWriterMock.ReadNexusFunc: method is nil but ConfigReadWriter.ReadNexus was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadNexus.Lock()
	mock.calls.ReadNexus = append(mock.calls.ReadNexus, callInfo)
	lockConfigReadWriterMockReadNexus.Unlock()
	return mock.ReadNexusFunc()
}

// ReadNexusCalls gets all the calls that were made to ReadNexus.
// Check the length with:
//     len(mockedConfigReadWriter.ReadNexusCalls())
func (mock *ConfigReadWriterMock) ReadNexusCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadNexus.RLock()
	calls = mock.calls.ReadNexus
	lockConfigReadWriterMockReadNexus.RUnlock()
	return calls
}

// ReadProduct calls ReadProductFunc.
func (mock *ConfigReadWriterMock) ReadProduct(product v1alpha1.ProductName) (ConfigReadable, error) {
	if mock.ReadProductFunc == nil {
		panic("ConfigReadWriterMock.ReadProductFunc: method is nil but ConfigReadWriter.ReadProduct was just called")
	}
	callInfo := struct {
		Product v1alpha1.ProductName
	}{
		Product: product,
	}
	lockConfigReadWriterMockReadProduct.Lock()
	mock.calls.ReadProduct = append(mock.calls.ReadProduct, callInfo)
	lockConfigReadWriterMockReadProduct.Unlock()
	return mock.ReadProductFunc(product)
}

// ReadProductCalls gets all the calls that were made to ReadProduct.
// Check the length with:
//     len(mockedConfigReadWriter.ReadProductCalls())
func (mock *ConfigReadWriterMock) ReadProductCalls() []struct {
	Product v1alpha1.ProductName
} {
	var calls []struct {
		Product v1alpha1.ProductName
	}
	lockConfigReadWriterMockReadProduct.RLock()
	calls = mock.calls.ReadProduct
	lockConfigReadWriterMockReadProduct.RUnlock()
	return calls
}

// ReadRHSSO calls ReadRHSSOFunc.
func (mock *ConfigReadWriterMock) ReadRHSSO() (*RHSSO, error) {
	if mock.ReadRHSSOFunc == nil {
		panic("ConfigReadWriterMock.ReadRHSSOFunc: method is nil but ConfigReadWriter.ReadRHSSO was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadRHSSO.Lock()
	mock.calls.ReadRHSSO = append(mock.calls.ReadRHSSO, callInfo)
	lockConfigReadWriterMockReadRHSSO.Unlock()
	return mock.ReadRHSSOFunc()
}

// ReadRHSSOCalls gets all the calls that were made to ReadRHSSO.
// Check the length with:
//     len(mockedConfigReadWriter.ReadRHSSOCalls())
func (mock *ConfigReadWriterMock) ReadRHSSOCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadRHSSO.RLock()
	calls = mock.calls.ReadRHSSO
	lockConfigReadWriterMockReadRHSSO.RUnlock()
	return calls
}

// ReadSolutionExplorer calls ReadSolutionExplorerFunc.
func (mock *ConfigReadWriterMock) ReadSolutionExplorer() (*SolutionExplorer, error) {
	if mock.ReadSolutionExplorerFunc == nil {
		panic("ConfigReadWriterMock.ReadSolutionExplorerFunc: method is nil but ConfigReadWriter.ReadSolutionExplorer was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadSolutionExplorer.Lock()
	mock.calls.ReadSolutionExplorer = append(mock.calls.ReadSolutionExplorer, callInfo)
	lockConfigReadWriterMockReadSolutionExplorer.Unlock()
	return mock.ReadSolutionExplorerFunc()
}

// ReadSolutionExplorerCalls gets all the calls that were made to ReadSolutionExplorer.
// Check the length with:
//     len(mockedConfigReadWriter.ReadSolutionExplorerCalls())
func (mock *ConfigReadWriterMock) ReadSolutionExplorerCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadSolutionExplorer.RLock()
	calls = mock.calls.ReadSolutionExplorer
	lockConfigReadWriterMockReadSolutionExplorer.RUnlock()
	return calls
}

// ReadThreeScale calls ReadThreeScaleFunc.
func (mock *ConfigReadWriterMock) ReadThreeScale() (*ThreeScale, error) {
	if mock.ReadThreeScaleFunc == nil {
		panic("ConfigReadWriterMock.ReadThreeScaleFunc: method is nil but ConfigReadWriter.ReadThreeScale was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadThreeScale.Lock()
	mock.calls.ReadThreeScale = append(mock.calls.ReadThreeScale, callInfo)
	lockConfigReadWriterMockReadThreeScale.Unlock()
	return mock.ReadThreeScaleFunc()
}

// ReadThreeScaleCalls gets all the calls that were made to ReadThreeScale.
// Check the length with:
//     len(mockedConfigReadWriter.ReadThreeScaleCalls())
func (mock *ConfigReadWriterMock) ReadThreeScaleCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadThreeScale.RLock()
	calls = mock.calls.ReadThreeScale
	lockConfigReadWriterMockReadThreeScale.RUnlock()
	return calls
}

// WriteConfig calls WriteConfigFunc.
func (mock *ConfigReadWriterMock) WriteConfig(config ConfigReadable) error {
	if mock.WriteConfigFunc == nil {
		panic("ConfigReadWriterMock.WriteConfigFunc: method is nil but ConfigReadWriter.WriteConfig was just called")
	}
	callInfo := struct {
		Config ConfigReadable
	}{
		Config: config,
	}
	lockConfigReadWriterMockWriteConfig.Lock()
	mock.calls.WriteConfig = append(mock.calls.WriteConfig, callInfo)
	lockConfigReadWriterMockWriteConfig.Unlock()
	return mock.WriteConfigFunc(config)
}

// WriteConfigCalls gets all the calls that were made to WriteConfig.
// Check the length with:
//     len(mockedConfigReadWriter.WriteConfigCalls())
func (mock *ConfigReadWriterMock) WriteConfigCalls() []struct {
	Config ConfigReadable
} {
	var calls []struct {
		Config ConfigReadable
	}
	lockConfigReadWriterMockWriteConfig.RLock()
	calls = mock.calls.WriteConfig
	lockConfigReadWriterMockWriteConfig.RUnlock()
	return calls
}

// readConfigForProduct calls readConfigForProductFunc.
func (mock *ConfigReadWriterMock) readConfigForProduct(product v1alpha1.ProductName) (ProductConfig, error) {
	if mock.readConfigForProductFunc == nil {
		panic("ConfigReadWriterMock.readConfigForProductFunc: method is nil but ConfigReadWriter.readConfigForProduct was just called")
	}
	callInfo := struct {
		Product v1alpha1.ProductName
	}{
		Product: product,
	}
	lockConfigReadWriterMockreadConfigForProduct.Lock()
	mock.calls.readConfigForProduct = append(mock.calls.readConfigForProduct, callInfo)
	lockConfigReadWriterMockreadConfigForProduct.Unlock()
	return mock.readConfigForProductFunc(product)
}

// readConfigForProductCalls gets all the calls that were made to readConfigForProduct.
// Check the length with:
//     len(mockedConfigReadWriter.readConfigForProductCalls())
func (mock *ConfigReadWriterMock) readConfigForProductCalls() []struct {
	Product v1alpha1.ProductName
} {
	var calls []struct {
		Product v1alpha1.ProductName
	}
	lockConfigReadWriterMockreadConfigForProduct.RLock()
	calls = mock.calls.readConfigForProduct
	lockConfigReadWriterMockreadConfigForProduct.RUnlock()
	return calls
}
