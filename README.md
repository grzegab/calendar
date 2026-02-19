# Kalendarz

Do rezerwacji miejsc nauki jazdy. Widok  kalendarza typu "inifinity" na nastepne 5 miesiecy + 1 wstecz.
Kontrola przez websocket + vanilla js.

## System
- Request od frontu przechodzi przez wewentryz wall
- sprawdzanie czy user ma sesje i czy IP nie jest banowane (redis)
- wymiana websocket i na strone propozycja

## Problemy
logowanie typu oAuth zeby byla sesja. HTML generowany z GO. JS doklejony do html.
WebSocket zwraca jsona, ktory jest transfortmowany przez JS na HTML. 

## Dzialanie
User ma 30h na jazdy, chyba ze admin wybierze inaczej. Propozycje tlyko na 30h, pozniej jak odpanda propozycje to nowe dopiero moze dawac
Admin moze wszystko, zmieniac godziny itp.

Logowanie usera: tel + mail do rejestracji, pozniej kod na sms/mail (sprawdzanie bruteforca). Banowanie IP przy bruteforce.

## Profilowanie
Pprof do sprawdzenia profili, dodatkowo do monitorowania sentry, kibana albo inne gowno (ADR do ustalenia).

## Budowa app
Na prod wrzucany obra na argoCD z k8s. budowanie przez pipeline github.
