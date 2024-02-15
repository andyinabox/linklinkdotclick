import { FormBase } from './form-base'
class FormAddLink extends FormBase {
  constructor() {
    super()
    this.inputs.url.remove()
    this.querySelector('label').remove()
  }

  onSubmit() {
    this.broadcast('link-create-request')
  }
}
customElements.define('form-add-link', FormAddLink, { extends: 'form' })
