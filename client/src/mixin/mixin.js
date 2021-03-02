export default {
  methods: {
    showLoginSeccessMessage() {
      this.$buefy.toast.open({
        duration: 5000,
        message: 'ログインしました',
        type: 'is-success'
      })
    },
    showLoginFailMessage() {
      this.$buefy.toast.open({
        duration: 5000,
        message: 'ユーザー名かパスワード名が間違っています',
        type: 'is-danger'
      })
    },
    showSignupSuccessMessage() {
      this.$buefy.toast.open({
        duration: 5000,
        message: '新規登録完了しました',
        type: 'is-success'
      })
    },
    showSignupFailMessage() {
      this.$buefy.toast.open({
        duration: 5000,
        message: '新規登録に失敗しました',
        type: 'is-danger'
      })
    },
    showLogoutMessage() {
      this.$buefy.toast.open({
        duration: 5000,
        message: 'ログアウトしました',
        type: 'is-info'
      })
    },
    showDeleteMessage() {
      this.$buefy.toast.open({
        duration: 4000,
        message: '写真を削除しました',
        type: 'is-info'
      })
    }
  }
}
