<script lang="ts">
  let playerVisible = false;
  function togglePlayerCreate() {
    playerVisible = !playerVisible
  }
</script>

<link href='https://unpkg.com/boxicons@2.0.7/css/boxicons.min.css' rel='stylesheet'>
<div class="sidebar close">
    <ul class="nav-links">
      <li>
        <div class="iocn-link">
          <a href="/players">
            <i id="player" class='bx bx-user'></i>
            <!-- <span class="link_name">Player</span> -->
          </a>
          <!-- <i class='bx bxs-chevron-down arrow' ></i> -->
        </div>
        <!-- <ul class="sub-menu">
          <li><a class="link_name" href="/players">Player</a></li>
        </ul> -->
      </li>
      <li>
      </li>
      <li>
        <div class="iocn-link">
          <a href="/items">
            <i class='bx bx-book-alt' ></i>
            <!-- <span class="link_name">Items</span> -->
          </a>
          <!-- <i class='bx bxs-chevron-down arrow' ></i> -->
        </div>
        <!-- <ul class="sub-menu">
          <li><a class="link_name" href="/items">Items</a></li>
        </ul> -->
      </li>
      <li>
        <div class="iocn-link">
          <a href="/bag">
            <i class='bx bx-shopping-bag'></i>
          </a>
        </div>
      </li>
</ul>
  </div>
<style>
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Poppins', sans-serif;
}
.sidebar{
  position: fixed;
  top: 0px;
  left: 0px; 
  height: 100vh;
  width: 260px;
  background: #11101d;
  z-index: 100;
  transition: all 0.5s ease;
}
.sidebar.close{
  width: 78px;
}
.sidebar .logo-details i{
  font-size: 30px;
  color: #fff;
  height: 50px;
  min-width: 78px;
  text-align: center;
  line-height: 50px;
}
.sidebar .nav-links{
  height: 100%;
  padding: 30px 0 150px 0;
  overflow: auto;
}
.sidebar.close .nav-links{
  overflow: visible;
}
.sidebar .nav-links::-webkit-scrollbar{
  display: none;
}
.sidebar .nav-links li{
  position: relative;
  list-style: none;
  transition: all 0.4s ease;
}
.sidebar .nav-links li:hover{
  background: #1d1b31;
}
.sidebar .nav-links li .iocn-link{
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.sidebar.close .nav-links li .iocn-link{
  display: block
}
.sidebar .nav-links li i{
  height: 50px;
  min-width: 78px;
  text-align: center;
  line-height: 50px;
  color: #fff;
  font-size: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
}
.sidebar .nav-links li a{
  display: flex;
  align-items: center;
  text-decoration: none;
}

.sidebar .nav-links li .sub-menu a{
  color: #fff;
  font-size: 15px;
  padding: 5px 0;
  white-space: nowrap;
  opacity: 0.6;
  transition: all 0.3s ease;
}
.sidebar .nav-links li .sub-menu a:hover{
  opacity: 1;
}
.sidebar.close .profile-details i {
  display: none;
}


</style>