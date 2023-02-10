## Cephapi Portal

## Установка
```bash
git clone <this-repo>; cd <this-repo>
make release
cp config.toml.bak config.toml
vim config.toml # Настройка
cp systemd/cephapi.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable --now cephapi.service

# Alternative || Dev
make release
make prodrun # Запуск
```

