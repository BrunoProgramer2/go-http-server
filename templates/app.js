const URL_API = "https://nnypwfnqbyxfmkgfbckr.supabase.co"
const URL_API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im5ueXB3Zm5xYnl4Zm1rZ2ZiY2tyIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NjQxMjUyOTMsImV4cCI6MTk3OTcwMTI5M30.yfjXhoE9Fq0x9XN-NqQF_YnXasjUTGAWYYg3MBnbayY"

const supa = supabase.createClient(URL_API,URL_API_KEY)

user = []

document.getElementById("button").addEventListener("click",() => {
    const {err, session, userSupa} = supa.auth.signIn({
        "provider":"github"
    })
})
console.log(supa.auth.session())
