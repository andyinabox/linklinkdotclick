import './elements/form/state-manager'
import './elements/form/edit-mode-toggle'
import './elements/form/style-editor'
import './elements/li/link-element'
import './elements/ol/link-list'
import './elements/body/base-page'

// remove no-js class
document.querySelector('html').classList.remove('no-js')

// remove query
if (window.location.search) {
  const split = window.location.toString().split('?')
  history.replaceState({}, '', split[0])
}
