/*
 * File to stash local debug changes to. Never check this in with changes
 */

import {createRouterState} from './reducers/router'
import * as Tabs from './constants/tabs'
import {updateConfig} from './command-line.desktop.js'

let config = {
  allowStartupFailure: false,
  printRPC: false,
  showDevTools: false,
  showAllTrackers: false,
  reduxDevToolsEnable: false,
  redirectOnLogout: true,
  reduxDevToolsSelect: state => state, // only watch a subset of the store
  enableStoreLogging: false,
  enableActionLogging: true,
  forwardLogs: true,
  devStoreChangingFunctions: false,
  printOutstandingRPCs: false,
  reactPerf: false,
  overrideLoggedInTab: null,
  focusOnShow: true,
  printRoutes: false,
  skipLauncherDevtools: true,
  initialTabState: {},
  forceMainWindowPosition: null,
  closureStoreCheck: false,
  searchActive: false,
  logStatFrequency: 0,
  actionStatFrequency: 0,
}

if (__DEV__ && process.env.KEYBASE_LOCAL_DEBUG) {
  config.allowStartupFailure = true
  config.printRPC = true
  config.showDevTools = false
  config.showAllTrackers = false
  config.reduxDevToolsEnable = false
  config.redirectOnLogout = false
  config.reduxDevToolsSelect = state => state.tracker
  config.enableStoreLogging = true
  config.enableActionLogging = false
  config.forwardLogs = false
  config.devStoreChangingFunctions = true
  config.printOutstandingRPCs = true
  config.reactPerf = false
  config.overrideLoggedInTab = Tabs.settingsTab
  config.focusOnShow = false
  config.printRoutes = true
  config.initialTabState = {
    [Tabs.loginTab]: [],
    [Tabs.settingsTab]: ['devMenu', 'dumbSheet'],
  }
  config.logStatFrequency = 0.8
  config.actionStatFrequency = 0.8

  let envJson = null
  if (process.env.KEYBASE_LOCAL_DEBUG_JSON) {
    try {
      envJson = JSON.parse(process.env.KEYBASE_LOCAL_DEBUG_JSON)
    } catch (e) {
      console.warn('Invalid KEYBASE_LOCAL_DEBUG_JSON:', e)
    }
  }

  config = {...config, ...envJson}
}

config = updateConfig(config)

export const {
  enableActionLogging,
  allowStartupFailure,
  printRPC,
  showDevTools,
  showAllTrackers,
  reduxDevToolsSelect,
  reduxDevToolsEnable,
  enableStoreLogging,
  forwardLogs,
  devStoreChangingFunctions,
  printOutstandingRPCs,
  reactPerf,
  overrideLoggedInTab,
  focusOnShow,
  printRoutes,
  skipLauncherDevtools,
  forceMainWindowPosition,
  closureStoreCheck,
  searchActive,
  logStatFrequency,
  actionStatFrequency,
} = config

export function initTabbedRouterState (state) {
  if (!__DEV__ || !process.env.KEYBASE_LOCAL_DEBUG) {
    return state
  }

  const ts = config.initialTabState
  const tabState = {}
  Object.keys(ts).forEach(tab => { tabState[tab] = createRouterState(ts[tab], []) })

  return {
    ...state,
    tabs: {
      ...state.tabs,
      ...tabState,
    },
  }
}
