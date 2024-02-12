export class ProgressiveBody extends HTMLBodyElement {
  set loading(bool) {
    if (bool) {
      this.classList.add('loading')
    } else {
      this.classList.remove('loading')
    }
  }
  get loading() {
    return this.classList.contains('loading')
  }
}
