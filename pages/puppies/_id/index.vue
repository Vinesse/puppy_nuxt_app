<template>
    <section class="single-puppy">
        <h1>{{ puppy.title }}</h1>
        <div>
           <img :src="puppy.thumbnail" :alt="puppy.title">
        </div>
        <p>{{ puppy.descriptionText }}</p>
    </section>
</template>

<script>
export default {
    asyncData(context) {
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                //here we are setting a timeout to fake an async request
                //that would normally be made to the server
                resolve({
                    puppy: [
                        {
                            id: "1",
                            title: "Pug",
                            thumbnail: "http://cdn-www.dailypuppy.com/media/dogs/anonymous/8684/2008091699067_100_2824.jpg_w450.jpg",
                            previewText: "Really cute Pug",
                            descriptionText: "Very truly a really cute pug"
                        },
                        {
                            id: "2",
                            title: "Boston Terrier",
                            thumbnail: "http://www.puppy-4-sale.net/black-boston-terrier-puppy.jpg",
                            previewText: "Really cute Boston Terrier",
                            descriptionText: "Very truly a really cute Boston Terrier"
                        }
                    ].find(el => el.id === context.params.id)
                })
            }, 1500)
        })
    }
}
</script>


<style scoped>
.single-puppy {
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    padding: 30px;
    flex-flow: column;
}
</style>
