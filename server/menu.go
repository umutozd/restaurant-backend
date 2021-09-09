package server

import "net/http"

func (s *server) ListMenu(w http.ResponseWriter, r *http.Request) {}

func (s *server) CreateMenuItem(w http.ResponseWriter, r *http.Request) {}

func (s *server) ListMenuItems(w http.ResponseWriter, r *http.Request) {}

func (s *server) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {}

func (s *server) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {}
