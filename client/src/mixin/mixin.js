export default {
  methods: {
    showLoginSeccessMessage() {
      this.$buefy.toast.open({
        message: 'ログインしました',
        type: 'is-success'
      })
    },
    showLoginFailMessage() {
      this.$buefy.toast.open({
        message: 'ユーザー名かパスワード名が間違っています',
        type: 'is-danger'
      })
    },
    showSignupSuccessMessage() {
      this.$buefy.toast.open({
        message: '新規登録完了しました',
        type: 'is-success'
      })
    },
    showSignupFailMessage() {
      this.$buefy.toast.open({
        message: '新規登録に失敗しました',
        type: 'is-danger'
      })
    },
    showLogoutMessage() {
      this.$buefy.toast.open({
        message: 'ログアウトしました',
        type: 'is-info'
      })
    }
  }
}
