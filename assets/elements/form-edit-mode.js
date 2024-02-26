import { FormBase } from './form-base'
class FormEditMode extends FormBase {
  constructor() {
    super()
    this.btn = this.querySelector('button')
  }
  get isEditing() {
    return this._editing
  }
  set isEditing(v) {
    if (v) {
      this.btn.innerText = 'Done'
      this._editing = true
      this.broadcast('edit-mode-start')
    } else {
      this.btn.innerText = 'Edit'
      this._editing = false
      this.broadcast('edit-mode-stop')
    }
  }
  onSubmit() {
    this.isEditing = !this.isEditing
  }
}
customElements.define('form-edit-mode', FormEditMode, { extends: 'form' })
