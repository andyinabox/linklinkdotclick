import { eventsMixin } from '../lib/mixins'
import { ProgressiveForm } from './progressive-form'
class AddLinkForm extends ProgressiveForm {
  constructor() {
    super()
    this.inputs.url.remove()
    this.querySelector('label').remove()
  }

  onSubmit() {
    this.broadcast('link-create-request')
  }
}
Object.assign(AddLinkForm.prototype, eventsMixin)
customElements.define('add-link-form', AddLinkForm, { extends: 'form' })
