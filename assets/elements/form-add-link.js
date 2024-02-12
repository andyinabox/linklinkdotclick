import { ProgressiveForm } from '../lib/progressive-form'
class FormAddLink extends ProgressiveForm {
  constructor() {
    super()
    console.log('form add link')
  }
}
customElements.define('linky-add-link', FormAddLink, { extends: 'form' })
