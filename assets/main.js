import './components/elements/form-state'
import './components/elements/form-edit-mode'
import './components/elements/form-style-editor'
import './components/elements/link-list-item'
import './components/elements/link-list'
import './components/layouts/page-home'
import './components/layouts/page-info'
import './components/layouts/page-about'

// remove no-js class
document.querySelector('html').classList.remove('no-js')

// remove query
if (window.location.search) {
  const split = window.location.toString().split('?')
  history.replaceState({}, '', split[0])
}
