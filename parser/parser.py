from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By

driver = webdriver.Firefox()

url = "https://travel.yandex.ru/hotels/fatih/?adults=2&bbox=28.91140136675%2C40.974181396984434~28.99663183925%2C41.054405921015565&checkinDate=2024-10-22&checkoutDate=2024-10-30&childrenAges=&geoId=115707&navigationToken=25&oneNightChecked=false&searchPagePollingId=b32a5c74444f9debbd4d05556fa75dab-1-newsearch&selectedSortId=relevant-first"
driver.get(url)
assert "Python" in driver.title
elem = driver.find_element(By.NAME, "q")
elem.clear()
elem.send_keys("pycon")
elem.send_keys(Keys.RETURN)
assert "No results found." not in driver.page_source
driver.close()
