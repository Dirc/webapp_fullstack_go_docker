# Idea for blog

- Intro: Ready to learn some Go Lang
  - (After Games With Go)
  - Wanted to create a simple webapp and found this excelent blog.
  - I got exited and decided to spice things up a bit by dockerise this webapp and database setup.
  - Which was surprisingly easy!

- Tool of choice: VS Code. Since it has a nice Go plugin.

- Steps:
  - Follow blog. It starts with a simple `Hello World`. Note that this example is still available at the `/hello` uri. Hence we can use it to verify if the webapp is running properly in the container.
  - Get webapp runnning in Go Docker
  - Get database running in Postgres Docker
  - Connect both using Docker Compose
  - Ensure webapp is using the Postgres container as database.
