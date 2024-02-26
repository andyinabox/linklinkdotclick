import './elements/form-state'
import './elements/form-edit-mode'
import './elements/form-style-editor'
import './elements/link-list-item'
import './elements/link-list'
import './elements/body-page'

// remove no-js class
document.querySelector('html').classList.remove('no-js')

// remove query
if (window.location.search) {
  const split = window.location.toString().split('?')
  history.replaceState({}, '', split[0])
}
