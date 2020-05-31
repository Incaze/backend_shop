<template>
    <div class="v-auth-info">
        <div v-if="this.AUTH_TEXT_STATUS.length" class="v-registration__status info_bg">
            <p class="info_text">{{this.AUTH_TEXT_STATUS}}</p>
        </div>
    </div>
</template>

<script>
    import {mapGetters} from "vuex";

    export default {
        name: "v-auth-info",
        computed: {
            ...mapGetters([
                'AUTH_TEXT_STATUS'
            ])
        }
    }
</script>

<style scoped>

</style>